package schema_test

import (
	"fmt"

	// mssql
	_ "github.com/denisenkom/go-mssqldb" // DriverName: mssql
	// gofreetds "github.com/minus5/gofreetds"      // DriverName: mssql

	// mysql
	_ "github.com/go-sql-driver/mysql"  // DriverName: mysql
	_ "github.com/ziutek/mymysql/godrv" // DriverName: mymysql

	// oracle
	// _ "gopkg.in/goracle.v2" // DriverName: goracle
	// _ "github.com/mattn/go-oci8" // DriverName: oci8
	// _ "gopkg.in/rana/ora.v4" // DriverName: ora

	// postgres
	_ "github.com/jackc/pgx/stdlib" // DriverName: pgx
	_ "github.com/lib/pq"           // DriverName: pg

	// _ "github.com/jbarham/gopgsqldriver" // DriverName: postgres

	// sqlite
	_ "github.com/mattn/go-sqlite3" // DriverName: sqlite3
	// _ "github.com/gwenn/gosqlite" // DriverName: sqlite3
	// _ "github.com/mxk/go-sqlite/sqlite3" // DriverName: sqlite3

	. "github.com/onsi/ginkgo"
	// . "github.com/onsi/gomega"
)

var _ = Describe("schema", func() {

	Context("using Microsoft SQL-Server", func() {
		const (
			user = "mssql_test_user"
			pass = "Password-123"
			host = "localhost"
			port = "41433"
		)

		Context("using driver github.com/denisenkom/go-mssqldb", func() {
			var params = mssqlDialect
			params.DriverName = "mssql"
			params.ConnStr = fmt.Sprintf("user id=%s;password=%s;server=%s;port=%s", user, pass, host, port)
			SchemaTestRunner(&params)
		})

		// TODO(js) Reinstate this. How do can we deal with duplicate driver names/ids?

		// Context("using driver github.com/minus5/gofreetds", func() {

		// 	// TODO(js) This results in:
		// 	// panic: sql: Register called twice for driver mssql
		// 	sql.Register("gofreetds", &gofreetds.MssqlDriver{})

		// 	var params = mssqlDialect
		// 	params.DriverName = "gofreetds"
		// 	params.ConnStr = fmt.Sprintf("user id=%s;password=%s;server=%s:%s", user, pass, host, port)
		// 	SchemaTestRunner(&params)
		// })
	})

	Context("using MySQL", func() {
		const (
			user = "mysql_test_user"
			pass = "password"
			host = "localhost"
			port = "43306"
			dbs  = "test_db"
		)

		Context("using github.com/go-sql-driver/mysql", func() {
			var params = mysqlDialect
			params.DriverName = "mysql"
			params.ConnStr = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, pass, host, port, dbs)
			SchemaTestRunner(&params)
		})
		Context("using github.com/ziutek/mymysql/godrv", func() {
			var params = mysqlDialect
			params.DriverName = "mymysql"
			params.ConnStr = fmt.Sprintf("tcp:%s:%s*%s/%s/%s", host, port, dbs, user, pass)
			SchemaTestRunner(&params)
		})
	})

	XContext("using Oracle", func() {
		const (
			user = "test_user"
			pass = "password"
			host = "localhost"
			port = "32772"
			dbs  = "xe"
		)

		Context("using github.com/go-goracle/goracle", func() {

			var params = oracleDialect
			params.DriverName = "goracle"
			// params.DriverName = "oci8"
			// params.DriverName = "ora"
			params.ConnStr = fmt.Sprintf("%s/%s@%s:%s/%s", user, pass, host, port, dbs)

			SchemaTestRunner(&params)
		})
	})

	Context("using Postgres", func() {
		const (
			user = "postgres"
			host = "localhost"
			port = "45432"
		)

		Context("using github.com/lib/pq", func() {
			var params = postgresDialect
			params.DriverName = "postgres"
			// params.DriverName = "pgx"
			params.ConnStr = fmt.Sprintf("user=%s host=%s port=%s sslmode=disable", user, host, port)
			SchemaTestRunner(&params)
		})

		Context("using github.com/jackc/pgx/stdlib", func() {
			var params = postgresDialect
			params.DriverName = "pgx"
			params.ConnStr = fmt.Sprintf("user=%s host=%s port=%s sslmode=disable", user, host, port)
			SchemaTestRunner(&params)
		})

		// TODO(js) Reinstate this test - but there's a name clash
		// - put test in seperate package?
		// - does import ordering force the execution order of init funcs?

		// Context("using github.com/jbarham/gopgsqldriver", func() {

		// 	sql.Register("pgsqldriver", &gopgsqldriver.postgresDriver{}) // TODO(js) It's private :/

		// 	var params = postgresDialect
		// 	params.DriverName = "gopgsqldriver"
		// 	params.ConnStr = fmt.Sprintf("user=%s host=%s port=%s sslmode=disable", user, host, port)
		// 	SchemaTestRunner(&params)
		// })
	})

	Context("using SQLite", func() {
		const (
			// user = ""
			// pass = ""
			// host = ""
			// port = ""
			dbs = ":memory:"
			// dbs = "./test.sqlite"
		)
		Context("using github.com/mattn/go-sqlite3", func() {

			// TODO(js) I think dbs connect string needs uniquing for SQLite?

			var params = sqliteDialect
			params.DriverName = "sqlite3"
			params.ConnStr = dbs

			SchemaTestRunner(&params)
		})
	})
})
