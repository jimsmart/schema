package schema_test

import (
	"fmt"

	_ "github.com/lib/pq" // postgres
	// _ "github.com/jackc/pgx/stdlib" // pgx
	// _ "github.com/jbarham/gopgsqldriver" // postgres

	. "github.com/onsi/ginkgo"
	// . "github.com/onsi/gomega"
)

var _ = XDescribe("schema", func() {
	Context("using github.com/lib/pq (Postgres)", func() {

		const (
			user = "postgres"
			// pass = ""
			host = "localhost"
			port = "32770"
			dbs  = "postgres"
		)

		var postgres = &testParams{
			DriverName: "postgres",
			// DriverName: "pgx",
			ConnStr: fmt.Sprintf("user=%s host=%s port=%s dbname=%s sslmode=disable", user, host, port, dbs),

			CreateDDL: []string{`
				CREATE TABLE IF NOT EXISTS web_resource (
					id				INTEGER NOT NULL,
					url				TEXT NOT NULL UNIQUE,
					content			BYTEA,
					compressed_size	INTEGER NOT NULL,
					content_length	INTEGER NOT NULL,
					content_type	TEXT NOT NULL,
					etag			TEXT NOT NULL,
					last_modified	TEXT NOT NULL,
					created_at		TIMESTAMP NOT NULL,
					modified_at		TIMESTAMP,
					PRIMARY KEY (id)
				)`,
				`CREATE INDEX IF NOT EXISTS idx_web_resource_url ON web_resource(url)`,
				`CREATE INDEX IF NOT EXISTS idx_web_resource_created_at ON web_resource(created_at)`,
				`CREATE INDEX IF NOT EXISTS idx_web_resource_modified_at ON web_resource(modified_at)`,
				`CREATE VIEW web_resource_view AS SELECT id, url FROM web_resource`,
			},
			DropDDL: []string{
				`DROP VIEW IF EXISTS web_resource_view`,
				`DROP INDEX IF EXISTS idx_web_resource_modified_at`,
				`DROP INDEX IF EXISTS idx_web_resource_created_at`,
				`DROP INDEX IF EXISTS idx_web_resource_url`,
				`DROP TABLE IF EXISTS web_resource`,
			},

			TableExpRes: []string{
				"id INT4",
				"url TEXT",
				"content BYTEA",
				"compressed_size INT4",
				"content_length INT4",
				"content_type TEXT",
				"etag TEXT",
				"last_modified TEXT",
				"created_at TIMESTAMP",
				"modified_at TIMESTAMP",
			},
			ViewExpRes: []string{
				"id INT4",
				"url TEXT",
			},

			TableNameExpRes: "web_resource",
			ViewNameExpRes:  "web_resource_view",
		}

		SchemaTestRunner(postgres)
	})
})
