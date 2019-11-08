package schema_test

import (
	"fmt"

	_ "github.com/denisenkom/go-mssqldb" // mssql
	// _ "github.com/minus5/gofreetds" // mssql

	. "github.com/onsi/ginkgo"
	// . "github.com/onsi/gomega"
)

// Setup of test db is performed on Docker startup
// via script: docker-db-init-mssql.sql

var _ = Describe("schema", func() {
	Context("using github.com/denisenkom/go-mssqldb (Microsoft SQL-Server)", func() {

		const (
			user = "mssql_test_user"
			pass = "Password-123"
			host = "localhost"
			port = "41433"
		)

		var mssql = &testParams{
			DriverName: "mssql",
			ConnStr:    fmt.Sprintf("user id=%s;password=%s;server=%s;port=%s", user, pass, host, port),
			// ConnStr: fmt.Sprintf("user id=%s;password=%s;server=%s:%s", user, pass, host, port), // gofreetds

			CreateDDL: []string{
				`CREATE TABLE web_resource (
					id				INTEGER NOT NULL,
					url				NVARCHAR NOT NULL UNIQUE,
					content			VARBINARY,
					compressed_size	INTEGER NOT NULL,
					content_length	INTEGER NOT NULL,
					content_type	NVARCHAR NOT NULL,
					etag			NVARCHAR NOT NULL,
					last_modified	NVARCHAR NOT NULL,
					created_at		DATETIME NOT NULL,
					modified_at		DATETIME,
					PRIMARY KEY (id)
				)`,
				`CREATE INDEX idx_web_resource_url ON web_resource(url)`,
				`CREATE INDEX idx_web_resource_created_at ON web_resource (created_at)`,
				`CREATE INDEX idx_web_resource_modified_at ON web_resource (modified_at)`,
				`CREATE VIEW web_resource_view AS SELECT id, url FROM web_resource`, // TODO gofreetds barfs on this!?
				`CREATE TABLE [blanks in name] (
					id INTEGER NOT NULL,
					PRIMARY KEY (id)
				)`,
			},
			DropDDL: []string{
				`DROP TABLE [blanks in name]`,
				`DROP VIEW IF EXISTS web_resource_view`,
				`DROP INDEX IF EXISTS idx_web_resource_modified_at ON web_resource`,
				`DROP INDEX IF EXISTS idx_web_resource_created_at ON web_resource`,
				`DROP INDEX IF EXISTS idx_web_resource_url ON web_resource`,
				`DROP TABLE web_resource`,
			},

			TableExpRes: []string{
				"id INT",
				"url NVARCHAR",
				"content VARBINARY",
				"compressed_size INT",
				"content_length INT",
				"content_type NVARCHAR",
				"etag NVARCHAR",
				"last_modified NVARCHAR",
				"created_at DATETIME",
				"modified_at DATETIME",
			},
			ViewExpRes: []string{
				"id INT",
				"url NVARCHAR",
			},

			TableNamesExpRes: []string{
				"web_resource",
				"blanks in name",
				// "[brackets] in name",
				// `"d.quotes" in name`,
				// "'s.quotes' in name",
				// "{braces} in name",
				// "`backticks` in name",
			},
			ViewNamesExpRes: []string{
				"web_resource_view",
			},
		}

		SchemaTestRunner(mssql)
	})
})
