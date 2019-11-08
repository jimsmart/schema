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
			// user = ""
			// pass = ""
			// host = ""
			// port = ""
			dbs = ":memory:"
			// dbs = "./test.sqlite"
		)

		var sqlite = &testParams{
			DriverName: "sqlite3",
			ConnStr:    dbs,

			CreateDDL: []string{
				`CREATE TABLE IF NOT EXISTS web_resource (
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
				"CREATE INDEX IF NOT EXISTS idx_web_resource_url ON web_resource(url)",
				"CREATE INDEX IF NOT EXISTS idx_web_resource_created_at ON web_resource(created_at)",
				"CREATE INDEX IF NOT EXISTS idx_web_resource_modified_at ON web_resource(modified_at)",
				"CREATE VIEW web_resource_view AS SELECT id, url FROM web_resource",
				// Tests for correct identifer escaping.
				`CREATE TABLE IF NOT EXISTS "blanks in name" (id INTEGER, PRIMARY KEY (id))`,
				`CREATE TABLE "[brackets] in name" (id INTEGER, PRIMARY KEY (id))`,
				`CREATE TABLE """d.quotes"" in name" (id INTEGER, PRIMARY KEY (id))`,
				`CREATE TABLE "'s.quotes' in name" (id INTEGER, PRIMARY KEY (id))`,
				`CREATE TABLE "{braces} in name" (id INTEGER, PRIMARY KEY (id))`,
				"CREATE TABLE \"`backticks` in name\" (id INTEGER, PRIMARY KEY (id))",
			},
			DropDDL: []string{
				"DROP TABLE \"`backticks` in name\"",
				`DROP TABLE "{braces} in name"`,
				`DROP TABLE "'s.quotes' in name"`,
				`DROP TABLE """d.quotes"" in name"`,
				`DROP TABLE "[brackets] in name"`,
				`DROP TABLE "blanks in name"`,
				"DROP VIEW web_resource_view",
				"DROP INDEX idx_web_resource_modified_at",
				"DROP INDEX idx_web_resource_created_at",
				"DROP INDEX idx_web_resource_url",
				"DROP TABLE web_resource",
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
				"id INTEGER",
				"url TEXT",
				"content BLOB",
				"compressed_size INTEGER",
				"content_length INTEGER",
				"content_type TEXT",
				"etag TEXT",
				"last_modified TEXT",
				"created_at DATETIME",
				"modified_at DATETIME",
			},
			ViewExpRes: []string{
				"id INTEGER",
				"url TEXT",
			},

			TableNamesExpRes: []string{
				"web_resource",
				// Tests for correct identifer escaping.
				"blanks in name",
				"[brackets] in name",
				`"d.quotes" in name`,
				"'s.quotes' in name",
				"{braces} in name",
				"`backticks` in name",
			},
			ViewNamesExpRes: []string{
				"web_resource_view",
			},
		}

		SchemaTestRunner(sqlite)
	})
})
