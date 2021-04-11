package schema_test

import (
	"fmt"

	_ "github.com/lib/pq" // postgres
	// _ "github.com/jackc/pgx/stdlib" // pgx
	// _ "github.com/jbarham/gopgsqldriver" // postgres

	. "github.com/onsi/ginkgo"
	// . "github.com/onsi/gomega"
)

// Database/user setup not needed: default database and schema are empty on Postgres.

var _ = Describe("schema", func() {
	Context("using github.com/lib/pq (Postgres)", func() {

		const (
			user = "postgres"
			host = "localhost"
			port = "45432"
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
				`CREATE TABLE person (
					given_name		TEXT NOT NULL,
					family_name		TEXT NOT NULL,
					PRIMARY KEY (family_name, given_name)
				)`,
			},
			DropDDL: []string{
				`DROP TABLE person`,
				`DROP VIEW web_resource_view`,
				`DROP INDEX idx_web_resource_modified_at`,
				`DROP INDEX idx_web_resource_created_at`,
				`DROP INDEX idx_web_resource_url`,
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

			TableNamesExpRes: [][2]string{
				{"public", "person"},
				{"public", "web_resource"},
			},
			ViewNamesExpRes: [][2]string{
				{"public", "web_resource_view"},
			},

			PrimaryKeysExpRes: []string{"family_name", "given_name"},
		}

		SchemaTestRunner(postgres)
	})
})
