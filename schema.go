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
func fetchColumnTypes(db *sql.DB, query, name string, escapeIdent func(string) string) ([]*sql.ColumnType, error) {
	query = fmt.Sprintf(query, escapeIdent(name))
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return rows.ColumnTypes()
}

// Tables returns column type metadata for all tables in the current schema.
// The returned map is keyed by table name.
func Tables(db *sql.DB) (map[string][]*sql.ColumnType, error) {
	d, err := getDialect(db)
	if err != nil {
		return nil, err
	}
	names, err := d.TableNames(db)
	if err != nil {
		return nil, err
	}
	if len(names) == 0 {
		return nil, nil
	}
	m := make(map[string][]*sql.ColumnType, len(names))
	for _, n := range names {
		ct, err := d.Table(db, n)
		if err != nil {
			return nil, err
		}
		m[n] = ct
	}
	return m, nil
}

// Views returns column type metadata for all views in the current schema.
// The returned map is keyed by view name.
func Views(db *sql.DB) (map[string][]*sql.ColumnType, error) {
	d, err := getDialect(db)
	if err != nil {
		return nil, err
	}
	names, err := d.ViewNames(db)
	if err != nil {
		return nil, err
	}
	if len(names) == 0 {
		return nil, nil
	}
	m := make(map[string][]*sql.ColumnType, len(names))
	for _, n := range names {
		ct, err := d.View(db, n)
		if err != nil {
			return nil, err
		}
		m[n] = ct
	}
	return m, nil
}
