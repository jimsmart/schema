package schema_test

import (
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3" // sqlite3
	// _ "github.com/gwenn/gosqlite" // sqlite3
	// _ "github.com/mxk/go-sqlite/sqlite3" // sqlite3

	. "github.com/onsi/ginkgo"
	// . "github.com/onsi/gomega"
)

var _ = Describe("schema", func() {
	Context("using github.com/mattn/go-sqlite3 (SQLite)", func() {

		const (
			dbs = ":memory:"
			// dbs = "./test.sqlite"
		)

		var sqlite = &testParams{
			DriverName: "sqlite3",
			ConnStr:    dbs,

			CreateDDL: []string{`
				CREATE TABLE IF NOT EXISTS web_resource (
					id				INTEGER NOT NULL,
					url				TEXT NOT NULL UNIQUE,
					content			BLOB,
					compressed_size	INTEGER NOT NULL,
					content_length	INTEGER NOT NULL,
					content_type	TEXT NOT NULL,
					etag			TEXT NOT NULL,
					last_modified	TEXT NOT NULL,
					created_at		DATETIME NOT NULL,
					modified_at		DATETIME,
					PRIMARY KEY (id)
				)`,
				`CREATE INDEX IF NOT EXISTS idx_web_resource_url ON web_resource(url)`,
				`CREATE INDEX IF NOT EXISTS idx_web_resource_created_at ON web_resource(created_at)`,
				`CREATE INDEX IF NOT EXISTS idx_web_resource_modified_at ON web_resource(modified_at)`,
				`CREATE VIEW web_resource_view AS SELECT id, url FROM web_resource`,
				`CREATE TABLE IF NOT EXISTS person (
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
			DropFn: func() {
				if dbs == ":memory:" {
					return
				}
				err := os.Remove(dbs)
				if err != nil {
					log.Printf("os.Remove error %v", err)
				}
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
				{"", "person"},
				{"", "web_resource"},
			},
			ViewNamesExpRes: [][2]string{
				{"", "web_resource_view"},
			},
			MaterializedViewNamesExpRes: [][2]string{},

			PrimaryKeysExpRes: []string{"family_name", "given_name"},
		}

		SchemaTestRunner(sqlite)
	})
})
