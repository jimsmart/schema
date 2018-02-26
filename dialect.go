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
	"*sqlite.impl":                &sqlite,   // github.com/gwenn/gosqlite
	"sqlite3.Driver":              &sqlite,   // github.com/mxk/go-sqlite - TODO(js) No datatypes.
	"*pq.Driver":                  &postgres, // github.com/lib/pq
	"*stdlib.Driver":              &postgres, // github.com/jackc/pgx
	"*pgsqldriver.postgresDriver": &postgres, // github.com/jbarham/gopgsqldriver - TODO(js) No datatypes.
	"*mysql.MySQLDriver":          &mysql,    // github.com/go-sql-driver/mysql
	"*godrv.Driver":               &mysql,    // github.com/ziutek/mymysql - TODO(js) No datatypes.
	"*mssql.MssqlDriver":          &mssql,    // github.com/denisenkom/go-mssqldb
	"*freetds.MssqlDriver":        &mssql,    // github.com/minus5/gofreetds - TODO(js) No datatypes. Error on create view.
	"*goracle.drv":                &oracle,   // gopkg.in/goracle.v2
	"*ora.Drv":                    &oracle,   // gopkg.in/rana/ora.v4 - TODO(js) Mismatched datatypes.
	"*oci8.OCI8Driver":            &oracle,   // github.com/mattn/go-oci8 - TODO(js) Mismatched datatypes.
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
