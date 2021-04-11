package schema_test

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql" // mysql
	// _ "github.com/ziutek/mymysql/godrv" // mymysql

	. "github.com/onsi/ginkgo"
	// . "github.com/onsi/gomega"
)

// Database/user setup by Docker in: docker-compose.yml

var _ = Describe("schema", func() {
	Context("using github.com/go-sql-driver/mysql (MySQL)", func() {

		const (
			user = "test_user"
			pass = "password-123"
			host = "localhost"
			port = "43306"
			dbs  = "test_db"
		)

		var mysql = &testParams{
			DriverName: "mysql",
			// DriverName: "mymysql",
			ConnStr: fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, pass, host, port, dbs), // mysql
			// ConnStr: fmt.Sprintf("tcp:%s:%s*%s/%s/%s", host, port, dbs, user, pass), // mymysql

			CreateDDL: []string{`
				CREATE TABLE IF NOT EXISTS web_resource (
					id				INTEGER NOT NULL,
					url				VARCHAR(255) NOT NULL UNIQUE, -- TODO(js) Earlier MySQL cannot handle UNIQUE with 1024 length.
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
				)`,
				`CREATE VIEW web_resource_view AS SELECT id, url FROM web_resource`,
				`CREATE TABLE IF NOT EXISTS person (
					given_name		VARCHAR(128) NOT NULL,
					family_name		VARCHAR(128) NOT NULL,
					PRIMARY KEY (family_name, given_name)
				)`,
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
				{"test_db", "person"},
				{"test_db", "web_resource"},
			},
			ViewNamesExpRes: [][2]string{
				{"test_db", "web_resource_view"},
			},

			PrimaryKeysExpRes: []string{"family_name", "given_name"},
		}

		SchemaTestRunner(mysql)
	})
})
