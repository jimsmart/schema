package schema_test

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	. "github.com/onsi/ginkgo"
	// . "github.com/onsi/gomega"
)

var _ = Describe("schema", func() {
	Context("using github.com/go-sql-driver/mysql (MySQL)", func() {

		const (
			user = "test_user"
			pass = "password"
			host = "localhost"
			port = "32778"
			dbs  = "test_db"
		)

		var mysql = &testParams{
			DriverName: "mysql",
			ConnStr:    fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, pass, host, port, dbs),

			CreateDDL: []string{`
				CREATE TABLE IF NOT EXISTS web_resource (
					id				INTEGER NOT NULL,
					url				VARCHAR(1024) NOT NULL UNIQUE,
					content			BLOB,
					compressed_size	INTEGER NOT NULL,
					content_length	INTEGER NOT NULL,
					content_type	VARCHAR(128) NOT NULL,
					etag			VARCHAR(128) NOT NULL,
					last_modified	VARCHAR(128) NOT NULL,
					created_at		TIMESTAMP NOT NULL,
					modified_at		TIMESTAMP NULL DEFAULT NULL,
					PRIMARY KEY (id),
					INDEX (url),
					INDEX (created_at),
					INDEX (modified_at)
				);`,
				`CREATE VIEW web_resource_view AS SELECT id, url FROM web_resource;`,
			},
			DropDDL: []string{
				`DROP VIEW web_resource_view;`,
				`DROP TABLE web_resource;`,
			},

			TableExpRes: []string{
				"id INT",
				"url VARCHAR",
				"content BLOB",
				"compressed_size INT",
				"content_length INT",
				"content_type VARCHAR",
				"etag VARCHAR",
				"last_modified VARCHAR",
				"created_at TIMESTAMP",
				"modified_at TIMESTAMP",
			},
			ViewExpRes: []string{
				"id INT",
				"url VARCHAR",
			},

			TableNameExpRes: "web_resource",
			ViewNameExpRes:  "web_resource_view",
		}

		SchemaTestRunner(mysql)
	})
})
