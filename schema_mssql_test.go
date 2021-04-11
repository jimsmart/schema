package schema_test

import (
	"fmt"

	_ "github.com/denisenkom/go-mssqldb" // mssql
	// _ "github.com/minus5/gofreetds" // mssql

	. "github.com/onsi/ginkgo"
	// . "github.com/onsi/gomega"
)

// Database/user setup script, run by Docker: docker-db-init-mssql.sql

var _ = Describe("schema", func() {
	Context("using github.com/denisenkom/go-mssqldb (Microsoft SQL-Server)", func() {

		const (
			// user = "mssql_test_user"
			user = "test_user"
			pass = "Password-123"
			host = "localhost"
			port = "41433"
		)

		var mssql = &testParams{
			DriverName: "mssql",
			ConnStr:    fmt.Sprintf("user id=%s;password=%s;server=%s;port=%s", user, pass, host, port),
			// ConnStr: fmt.Sprintf("user id=%s;password=%s;server=%s:%s", user, pass, host, port), // gofreetds

			CreateDDL: []string{`
				CREATE TABLE web_resource (
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
				// `CREATE VIEW web_resource_view AS SELECT t.id, t.url FROM web_resource t`,
				// `CREATE VIEW web_resource_view AS (SELECT id, url FROM web_resource)`,
				// `CREATE VIEW web_resource_view AS (SELECT t.id, t.url FROM web_resource t)`,
				`CREATE TABLE person (
					given_name		NVARCHAR NOT NULL,
					family_name		NVARCHAR NOT NULL,
					PRIMARY KEY (family_name, given_name)
				)`,
			},
			DropDDL: []string{
				`DROP TABLE person`,
				`DROP VIEW IF EXISTS web_resource_view`,
				`DROP INDEX IF EXISTS idx_web_resource_modified_at ON web_resource`,
				`DROP INDEX IF EXISTS idx_web_resource_created_at ON web_resource`,
				`DROP INDEX IF EXISTS idx_web_resource_url ON web_resource`,
				`DROP TABLE web_resource`,
			},

			TableExpRes: []string{
				"id",
				"url",
				"content",
				"compressed_size",
				"content_length",
				"content_type",
				"etag",
				"last_modified",
				"created_at",
				"modified_at",
			},
			ViewExpRes: []string{
				"id",
				"url",
			},

			TableNamesExpRes: []string{"person", "web_resource"},
			ViewNamesExpRes:  []string{"web_resource_view"},

			PrimaryKeysExpRes: []string{"family_name", "given_name"},

			TableNamesWithSchemaExpRes: [][2]string{
				{"test_db", "person"},
				{"test_db", "web_resource"},
			},

			ViewNamesWithSchemaExpRes: [][2]string{
				{"test_db", "web_resource_view"},
			},
		}

		SchemaTestRunner(mssql)
	})
})
