package schema_test

import (
	"fmt"

	_ "github.com/lib/pq" // postgres
	// _ "github.com/jackc/pgx/stdlib" // pgx
	// _ "github.com/jbarham/gopgsqldriver" // postgres

	. "github.com/onsi/ginkgo"
	// . "github.com/onsi/gomega"
)

// No setup needed, default database and schema are empty on Postgres.

var _ = Describe("schema", func() {
	Context("using github.com/lib/pq (Postgres)", func() {

		const (
			user = "postgres"
			host = "localhost"
			// port = "32774"
			port = "5432"
		)

		var postgres = &testParams{
			DriverName: "postgres",
			// DriverName: "pgx",
			ConnStr: fmt.Sprintf("user=%s host=%s port=%s sslmode=disable", user, host, port),

			CreateDDL: []string{`
				CREATE TABLE web_resource (
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
				`CREATE INDEX idx_web_resource_url ON web_resource(url)`,
				`CREATE INDEX idx_web_resource_created_at ON web_resource(created_at)`,
				`CREATE INDEX idx_web_resource_modified_at ON web_resource(modified_at)`,
				`CREATE VIEW web_resource_view AS SELECT id, url FROM web_resource`,
			},
			DropDDL: []string{
				`DROP VIEW web_resource_view`,
				`DROP INDEX idx_web_resource_modified_at`,
				`DROP INDEX idx_web_resource_created_at`,
				`DROP INDEX idx_web_resource_url`,
				`DROP TABLE web_resource`,
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
