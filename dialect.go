package schema

import "strings"

// query defines dialect query types.
type query int

// query type enum.
const (
	columnTypes query = iota // Index of query to get column type info.
	tableNames               // Index of query to get table names.
	viewNames                // Index of query to get view names.
)

// dialect describes how each database 'flavour' provides its metadata.
type dialect struct {
	// queries for fetching metadata: tableNames, viewNames, columnTypes.
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

// pack a string, normalising its whitespace.
func pack(s string) string {
	return strings.Join(strings.Fields(s), " ")
}
