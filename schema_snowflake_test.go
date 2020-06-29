package schema_test

import (
	"fmt"

	_ "github.com/snowflakedb/gosnowflake" // snowflake
	// _ "github.com/jackc/pgx/stdlib" // pgx
	// _ "github.com/jbarham/gopgsqldriver" // postgres

	. "github.com/onsi/ginkgo"
	// . "github.com/onsi/gomega"
)

// No setup needed, default database and schema are empty on Postgres.

var _ = XDescribe("schema", func() {
	Context("using github.com/gosnowflake/snowflake (Snowflake)", func() {

		const (
			user     = ""
			pass     = ""
			host     = ""
			database = ""
		)

		var snowflake = &testParams{
			DriverName: "snowflake",
			ConnStr:    fmt.Sprintf("%s:%s@%s/%s", user, pass, host, database),

			CreateDDL: []string{
				`USE SCHEMA TEST;`,
				`CREATE TABLE web_resource (
					id				INTEGER NOT NULL,
					url				TEXT NOT NULL,
					content			BINARY,
					compressed_size	INTEGER NOT NULL,
					content_length	INTEGER NOT NULL,
					content_type	TEXT NOT NULL,
					etag			TEXT NOT NULL,
					last_modified	TEXT NOT NULL,
					created_at		TIMESTAMP NOT NULL,
					modified_at		TIMESTAMP
				)`,
				`CREATE VIEW web_resource_view AS SELECT id, url FROM web_resource`,
			},
			DropDDL: []string{
				`DROP TABLE web_resource`,
				`DROP VIEW web_resource_view`,
			},

			TableExpRes: []string{
				"ID FIXED",
				"URL TEXT",
				"CONTENT BINARY",
				"COMPRESSED_SIZE FIXED",
				"CONTENT_LENGTH FIXED",
				"CONTENT_TYPE TEXT",
				"ETAG TEXT",
				"LAST_MODIFIED TEXT",
				"CREATED_AT TIMESTAMP_NTZ",
				"MODIFIED_AT TIMESTAMP_NTZ",
			},
			ViewExpRes: []string{
				"ID FIXED",
				"URL TEXT",
			},

			TableNameExpRes: "WEB_RESOURCE",
			ViewNameExpRes:  "WEB_RESOURCE_VIEW",
		}

		SchemaTestRunner(snowflake)
	})
})
