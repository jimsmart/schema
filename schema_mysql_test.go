package schema_test

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql" // mysql
	// _ "github.com/ziutek/mymysql/godrv" // mymysql

	. "github.com/onsi/ginkgo"
	// . "github.com/onsi/gomega"
)

// Setup script:
//
// CREATE DATABASE test_db;
// CREATE USER test_user IDENTIFIED BY 'password';
// GRANT ALL ON test_db.* TO 'test_user';

var _ = Describe("schema", func() {
	Context("using github.com/go-sql-driver/mysql (MySQL)", func() {

		const (
			user = "test_user"
			pass = "password"
			host = "localhost"
			port = "3306"
			// port = "32769"
			dbs = "test_db"
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
			},
			DropDDL: []string{
				`DROP VIEW web_resource_view`,
				`DROP TABLE web_resource`,
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
