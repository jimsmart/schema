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
//  	fmt.Println("Table:", tnames[i])
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
//  	fmt.Println("Column:", tcols[i].Name(), tcols[i].DatabaseTypeName())
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
// Underlying support for column type metadata is driver implementation specific and somewhat variable.
//
// View Metadata
//
// The same metadata can also be queried for views.
//
//  vnames, err := schema.ViewNames(db)
//  	...
//  vcols, err := schema.View(db, "monthly_sales_view")
//  	...
//
// Primary Key Metadata
//
// To obtain a list of columns making up the primary key for a given table:
//
//  pks, err := schema.PrimaryKey(db, "employee_tbl")
//  	...
//  // pks is []string
//  for i := range pks {
//  	fmt.Println("Primary Key:", pks[i])
//  }
//
//  // Output:
//  // Primary Key: employee_id
//
package schema

import (
	"database/sql"
	"fmt"
)

// https://github.com/golang/go/issues/7408
//
// https://github.com/golang/go/issues/7408#issuecomment-252046876
//
// If this package were to be part of database/sql, then the API would become like:-
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

// TableNames returns a list of all table names in the current schema.
func TableNames(db *sql.DB) ([]string, error) {
	d, err := getDialect(db)
	if err != nil {
		return nil, err
	}
	return d.TableNames(db)
}

// ViewNames returns a list of all view names in the current schema.
func ViewNames(db *sql.DB) ([]string, error) {
	d, err := getDialect(db)
	if err != nil {
		return nil, err
	}
	return d.ViewNames(db)
}

// PrimaryKey returns a list of column names making up the primary
// key for the given table name.
func PrimaryKey(db *sql.DB, name string) ([]string, error) {
	d, err := getDialect(db)
	if err != nil {
		return nil, err
	}
	return d.PrimaryKey(db, name)
}

// fetchNames executes the given query with an optional name parameter,
// and returns a list of table/view/column names.
//
// The name parameter (if not "") is passed as a parameter to db.Query.
func fetchNames(db *sql.DB, query, name string) ([]string, error) {
	var rows *sql.Rows
	var err error
	if len(name) > 0 {
		rows, err = db.Query(query, name)
	} else {
		rows, err = db.Query(query)
	}
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

func getDialect(db *sql.DB) (dialect, error) {
	dt := fmt.Sprintf("%T", db.Driver())
	d, ok := driverDialect[dt]
	if !ok {
		return nil, UnknownDriverError{Driver: dt}
	}
	return d, nil
}

// Table returns the column type metadata for the given table name.
func Table(db *sql.DB, name string) ([]*sql.ColumnType, error) {
	d, err := getDialect(db)
	if err != nil {
		return nil, err
	}
	return d.Table(db, name)
}

// View returns the column type metadata for the given view name.
func View(db *sql.DB, name string) ([]*sql.ColumnType, error) {
	d, err := getDialect(db)
	if err != nil {
		return nil, err
	}
	return d.View(db, name)
}

// fetchColumnTypes queries the database and returns column's type metadata
// for a single table or view.
func fetchColumnTypes(db *sql.DB, query string) ([]*sql.ColumnType, error) {
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return rows.ColumnTypes()
}

// Tables returns column type metadata for all tables in the current schema
// (not including system tables). The returned map is keyed by table name.
func Tables(db *sql.DB) (map[string][]*sql.ColumnType, error) {
	d, err := getDialect(db)
	if err != nil {
		return nil, err
	}
	return d.Tables(db)
}

// Views returns column type metadata for all views in the current schema
// (not including system views). The returned map is keyed by view name.
func Views(db *sql.DB) (map[string][]*sql.ColumnType, error) {
	d, err := getDialect(db)
	if err != nil {
		return nil, err
	}
	return d.Views(db)
}

// listFunc provides a list of names from the database.
type listFunc func(*sql.DB) ([]string, error)

type typeFunc func(db *sql.DB, name string) ([]*sql.ColumnType, error)

// fetchColumnTypesAll queries the database and returns metadata about the
// column types for all tables or all views.
//
// It uses the passed list provider function to obtain table/view names,
// and calls fetchColumnTypes() to fetch the column metadata for each name in the list.
func fetchColumnTypesAll(db *sql.DB, nameFn listFunc, typeFn typeFunc) (map[string][]*sql.ColumnType, error) {
	names, err := nameFn(db)
	if err != nil {
		return nil, err
	}
	if len(names) == 0 {
		return nil, nil
	}
	m := make(map[string][]*sql.ColumnType, len(names))
	for _, n := range names {
		ct, err := typeFn(db, n)
		if err != nil {
			return nil, err
		}
		m[n] = ct
	}
	return m, nil
}
