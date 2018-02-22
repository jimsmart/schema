// Package schema provides database schema metadata for database/sql drivers.
package schema

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
)

// https://github.com/golang/go/issues/7408
//
// https://github.com/golang/go/issues/7408#issuecomment-252046876
//
// Proposed API:-
//
//  SchemaNames(db *sql.DB) ([]string, error)
//  SchemaObject(db *sql.DB, name string) ([]sql.ColumnType, error)
//  Schema(db *sql.DB) (map[string][]sql.ColumnType, error)
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
// If this were to be part of database/sql, the API would become:-
//
//  func (db *DB) Table(name string) ([]*ColumnType, error)
//  func (db *DB) TableNames() ([]string, error)
//  func (db *DB) Tables() (map[string][]*ColumnType, error)
//  func (db *DB) View(name string) ([]*ColumnType, error)
//  func (db *DB) ViewNames() ([]string, error)
//  func (db *DB) Views() (map[string][]*ColumnType, error)
//

//

// query defines dialect query types.
type query int

// query type enum.
const (
	tableNames  query = iota // Index of query to get table names.
	viewNames                // Index of query to get view names.
	columnTypes              // Index of query to get column type info.
)

// dialect describes how each 'flavour' of database provides its metadata.
type dialect struct {
	// queries for fetching database schema metadata,
	// one per query type (tableNames, viewNames and columnTypes).
	queries [3]string
}

// driverDialect is a registry, mapping database/sql driver names to database dialects.
// This is somewhat fragile.
var driverDialect map[string]*dialect = map[string]*dialect{
	"*sqlite3.SQLiteDriver":       &sqlite,   // github.com/mattn/go-sqlite3
	"*sqlite.impl":                &sqlite,   // TODO(js) UNTESTED github.com/gwenn/gosqlite
	"sqlite3.Driver":              &sqlite,   // TODO(js) UNTESTED github.com/mxk/go-sqlite
	"*pq.Driver":                  &postgres, // github.com/lib/pq
	"*stdlib.Driver":              &postgres, // TODO(js) UNTESTED github.com/jackc/pgx/stdlib
	"*pgsqldriver.postgresDriver": &postgres, // TODO(js) UNTESTED github.com/jbarham/gopgsqldriver
	"*mysql.MySQLDriver":          &mysql,    // github.com/go-sql-driver/mysql
	"*godrv.Driver":               &mysql,    // TODO(js) UNTESTED github.com/ziutek/mymysql
	"*mssql.MssqlDriver":          &mssql,    // github.com/denisenkom/go-mssqldb
	"*freetds.MssqlDriver":        &mssql,    // TODO(js) UNTESTED github.com/minus5/gofreetds
	"*goracle.drv":                &oracle,   // gopkg.in/goracle.v2
	"*ora.Drv":                    &oracle,   // TODO(js) UNTESTED gopkg.in/rana/ora.v4
	"*oci8.OCI8Driver":            &oracle,   // TODO(js) UNTESTED github.com/mattn/go-oci8
}

// TODO Should we expose a method of registering a driver string/dialect in our registry?
// -- It would allow folk to work around the fragility. e.g.
//
// func Register(driver sql.Driver, d *Dialect) {}
//

// TableNames returns a list of all table names in the database
// (not including system tables).
//
// If the underlying driver does not support this feature,
// or is not a driver recognised by this package,
// then (nil, nil) is returned, and a log message is emitted.
func TableNames(db *sql.DB) ([]string, error) {
	// Originally called 'SchemaNames' in comment/proposal.
	return names(db, tableNames)
}

// ViewNames returns a list of all view names in the database
// (not including system views).
//
// If the underlying driver does not support this feature,
// or is not a driver recognised by this package,
// then (nil, nil) is returned, and a log message is emitted.
func ViewNames(db *sql.DB) ([]string, error) {
	// Originally called 'SchemaNames' in comment/proposal.
	return names(db, viewNames)
}

func names(db *sql.DB, qt query) ([]string, error) {
	// Originally called 'SchemaNames' in comment/proposal.
	dt := fmt.Sprintf("%T", db.Driver())
	d, ok := driverDialect[dt]
	if !ok {
		log.Printf("unknown db driver %s\n", dt)
		return nil, nil
	}
	q := d.queries[qt]
	rows, err := db.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

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

// Table returns the column type information for the given table.
//
// If the underlying driver does not support this feature,
// or is not a driver recognised by this package,
// then (nil, nil) is returned, and a log message is emitted.
func Table(db *sql.DB, name string) ([]*sql.ColumnType, error) {
	// Originally called 'SchemaObject' in comment/proposal.
	return object(db, name)
}

// View returns the column type information for the given view.
//
// If the underlying driver does not support this feature,
// or is not a driver recognised by this package,
// then (nil, nil) is returned, and a log message is emitted.
func View(db *sql.DB, name string) ([]*sql.ColumnType, error) {
	// Originally called 'SchemaObject' in comment/proposal.
	return object(db, name)
}

func object(db *sql.DB, name string) ([]*sql.ColumnType, error) {
	// Originally called 'SchemaObject' in comment/proposal.
	dt := fmt.Sprintf("%T", db.Driver())
	d, ok := driverDialect[dt]
	if !ok {
		log.Printf("unknown db driver %s\n", dt)
		return nil, nil
	}
	q := fmt.Sprintf(d.queries[columnTypes], name)
	rows, err := db.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return rows.ColumnTypes()
}

// Tables returns column type information for all tables in the database
// (not including system tables). The returned map is keyed by table name.
//
// If the underlying driver does not support this feature,
// or is not a driver recognised by this package,
// then (nil, nil) is returned, and a log message is emitted.
func Tables(db *sql.DB) (map[string][]*sql.ColumnType, error) {
	// Originally called 'Schema' in comment/proposal.
	return objects(db, TableNames)
}

// Views returns column type information for all views in the database
// (not including system views). The returned map is keyed by table name.
//
// If the underlying driver does not support this feature,
// or is not a driver recognised by this package,
// then (nil, nil) is returned, and a log message is emitted.
func Views(db *sql.DB) (map[string][]*sql.ColumnType, error) {
	// Originally called 'Schema' in comment/proposal.
	return objects(db, ViewNames)
}

func objects(db *sql.DB, nameFn func(*sql.DB) ([]string, error)) (map[string][]*sql.ColumnType, error) {
	// Originally called 'Schema' in comment/proposal.
	names, err := nameFn(db)
	if err != nil {
		return nil, err
	}
	m := make(map[string][]*sql.ColumnType, len(names))
	for _, n := range names {
		ci, err := object(db, n)
		if err != nil {
			return nil, err
		}
		m[n] = ci
	}
	return m, nil
}

// pack a string, normalising its whitespace.
func pack(s string) string {
	return strings.Join(strings.Fields(s), " ")
}
