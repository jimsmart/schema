// Package schema provides access to database schema metadata, for database/sql drivers.
//
// For further information about current driver support status, see https://github.com/jimsmart/schema
//
// Table Metadata
//
// The schema package works alongside database/sql and its underlying driver to provide schema metadata.
//
//  tnames, err := schema.TableNames(db)
//  	...
//  // tnames is []string
//  for i := range tnames {
//  	fmt.Printf("Table: %s\n", tnames[i])
//  }
//
//  // Output:
//  // Table: employee_tbl
//  // Table: department_tbl
//  // Table: sales_tbl
//
// Both user permissions and current database/schema effect table visibility.
//
// To query column type metadata for a single table, use schema.Table().
//
//  tcols, err := schema.Table(db, "employee_tbl")
//  	...
//  // tcols is []*sql.ColumnType
//  for i := range tcols {
//  	fmt.Printf("Column: %s %s\n", tcols[i].Name(), tcols[i].DatabaseTypeName())
//  }
//
//  // Output:
//  // Column: employee_id INTEGER
//  // Column: first_name TEXT
//  // Column: last_name TEXT
//  // Column: created_at TIMESTAMP
//
// To query table names and column type metadata for all tables, use schema.Tables().
//
// See also https://golang.org/pkg/database/sql/#ColumnType
//
// Underlying driver support for column type metadata is implementation specific and somewhat variable.
//
// View Metadata
//
// The same metadata can also be queried for views.
package schema

import (
	"database/sql"
	"fmt"
)

// https://github.com/golang/go/issues/7408
//
// https://github.com/golang/go/issues/7408#issuecomment-252046876
//
// Last proposed API:-
//
//  SchemaNames(db *sql.DB) ([]string, error)
//  SchemaObject(db *sql.DB, name string) ([]sql.ColumnType, error)
//  Schema(db *sql.DB) (map[string][]sql.ColumnType, error)
//
//
// After some refactoring, this is where it's at:-
//
//  TableNames(db *sql.DB) ([]string, error)
//  Table(db *sql.DB, name string) ([]*sql.ColumnType, error)
//  Tables(db *sql.DB) (map[string][]*sql.ColumnType, error)
//
//  ViewNames(db *sql.DB) ([]string, error)
//  View(db *sql.DB, name string) ([]*sql.ColumnType, error)
//  Views(db *sql.DB) (map[string][]*sql.ColumnType, error)
//
//
// If this package were to be part of database/sql, then the API would become:-
//
//  func (db *DB) Table(name string) ([]*ColumnType, error)
//  func (db *DB) TableNames() ([]string, error)
//  func (db *DB) Tables() (map[string][]*ColumnType, error)
//  func (db *DB) View(name string) ([]*ColumnType, error)
//  func (db *DB) ViewNames() ([]string, error)
//  func (db *DB) Views() (map[string][]*ColumnType, error)
//

//

// UnknownDriverError is returned when there is no matching
// database driver type name in the driverDialect table.
//
// Errors of this kind are caused by using an unsupported
// database driver/dialect, or if/when a database driver
// developer renames the type underlying calls to db.Driver().
type UnknownDriverError struct {
	Driver string
}

// Error returns a formatted string description.
func (e UnknownDriverError) Error() string {
	return fmt.Sprintf("unknown database driver %s", e.Driver)
}

//

// TableNames returns a list of all table names in the current schema
// (not including system tables).
func TableNames(db *sql.DB) ([]string, error) {
	return fetchNames(db, tableNames)
}

// ViewNames returns a list of all view names in the current schema
// (not including system views).
func ViewNames(db *sql.DB) ([]string, error) {
	return fetchNames(db, viewNames)
}

// fetchNames queries the database schema metadata and returns
// either a list of table or view names.
//
// It uses the database driver name and the passed query type
// to lookup the appropriate dialect and query.
func fetchNames(db *sql.DB, qt query) ([]string, error) {
	dt := fmt.Sprintf("%T", db.Driver())
	d, ok := driverDialect[dt]
	if !ok {
		// // TODO(js) This should return an error instead of nil (and should not log?)
		// log.Printf("unknown db driver %s\n", dt)
		return nil, UnknownDriverError{Driver: dt}
	}
	// Run the appropriate query from dialect:
	// this runs a query to fetch names of tables/views
	// from tables that contain db metadata.
	// It's different for every dialect.
	q := d.queries[qt]
	rows, err := db.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	// Scan result into list of names.
	var names []string
	n := ""
	for rows.Next() {
		err = rows.Scan(&n)
		if err != nil {
			return nil, err
		}
		names = append(names, n)
	}
	return names, nil
}

// Table returns the column type metadata for the given table name.
func Table(db *sql.DB, name string) ([]*sql.ColumnType, error) {
	return fetchColumnTypes(db, name)
}

// View returns the column type metadata for the given view name.
func View(db *sql.DB, name string) ([]*sql.ColumnType, error) {
	return fetchColumnTypes(db, name)
}

// fetchColumnTypes queries the database and returns column's type metadata
// for a single table or view.
//
// It uses the database driver name to look up the appropriate
// dialect, and the passed table/view name to build the query.
func fetchColumnTypes(db *sql.DB, name string) ([]*sql.ColumnType, error) {
	dt := fmt.Sprintf("%T", db.Driver())
	d, ok := driverDialect[dt]
	if !ok {
		// // TODO(js) This should return an error instead of nil (and should not log?)
		// log.Printf("unknown db driver %s\n", dt)
		return nil, UnknownDriverError{Driver: dt}
	}
	// Build and run the appropriate query from dialect:
	// this runs a query that returns no rows, and then
	// picks off the column type info.
	q := fmt.Sprintf(d.queries[columnTypes], d.escapeIdent(name))
	rows, err := db.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return rows.ColumnTypes()
}

// Tables returns column type metadata for all tables in the current schema
// (not including system tables). The returned map is keyed by table name.
func Tables(db *sql.DB) (map[string][]*sql.ColumnType, error) {
	return fetchColumnTypesAll(db, TableNames)
}

// Views returns column type metadata for all views in the current schema
// (not including system views). The returned map is keyed by view name.
func Views(db *sql.DB) (map[string][]*sql.ColumnType, error) {
	return fetchColumnTypesAll(db, ViewNames)
}

// listFn provides a list of names from the database.
type listFn func(*sql.DB) ([]string, error)

// fetchColumnTypesAll queries the database and returns metadata about the
// column types for all tables or all views.
//
// It uses the passed list provider function to obtain table/view names,
// and calls fetchColumnTypes() to fetch the column metadata for each name in the list.
func fetchColumnTypesAll(db *sql.DB, nameFn listFn) (map[string][]*sql.ColumnType, error) {
	names, err := nameFn(db)
	if err != nil {
		return nil, err
	}
	if len(names) == 0 {
		return nil, nil
	}
	m := make(map[string][]*sql.ColumnType, len(names))
	for _, n := range names {
		ct, err := fetchColumnTypes(db, n)
		if err != nil {
			return nil, err
		}
		m[n] = ct
	}
	return m, nil
}
