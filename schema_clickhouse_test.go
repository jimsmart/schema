package schema_test

import (
	"fmt"

	_ "github.com/ClickHouse/clickhouse-go/v2" // clickhouse
	// _ "github.com/jackc/pgx/stdlib" // pgx
	// _ "github.com/jbarham/gopgsqldriver" // postgres

	. "github.com/onsi/ginkgo"
	// . "github.com/onsi/gomega"
)

// Database/user setup not needed: default database and schema are empty on Postgres.

var _ = Describe("schema", func() {
	Context("using github.Com/ClickHouse/clickhouse-go/v2 (Clickhouse)", func() {

		const (
			user = "default"
			host = "localhost"
			port = "49000"
		)

		var clickhouse = &testParams{
			DriverName: "clickhouse",
			// DriverName: "pgx",
			ConnStr: fmt.Sprintf("clickhouse://%s@%s:%s/default", user, host, port),

			CreateDDL: []string{`
			CREATE TABLE web_resource (
    		id UInt32 NOT NULL,
    		url String NOT NULL,
    		content LowCardinality(Nullable(String)) NOT NULL,
    		compressed_size UInt32 NOT NULL,
    		content_length UInt32 NOT NULL,
    		content_type LowCardinality(String) NOT NULL,
    		etag LowCardinality(String) NOT NULL,
    		last_modified LowCardinality(String) NOT NULL,
    		created_at DateTime('UTC') NOT NULL,
    		modified_at Nullable(DateTime('UTC')),
    		PRIMARY KEY (id)
			) ENGINE = MergeTree()
			ORDER BY id;
			`,
				`CREATE VIEW web_resource_view AS SELECT id, url FROM web_resource`,
				`CREATE TABLE person (
    			given_name String NOT NULL,
    			family_name String NOT NULL,
    			event_date DateTime DEFAULT now()
				) ENGINE = MergeTree()
				ORDER BY (family_name, given_name);
				`,
			},
			DropDDL: []string{
				`DROP TABLE person`,
				`DROP VIEW web_resource_view`,
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
				{"default", "person"},
				{"default", "web_resource"},
			},
			ViewNamesExpRes: [][2]string{
				{"default", "web_resource_view"},
			},

			PrimaryKeysExpRes: []string{"family_name", "given_name"},
		}

		SchemaTestRunner(clickhouse)
	})
})
