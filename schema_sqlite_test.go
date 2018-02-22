package schema_test

import (
	"database/sql"
	"log"
	"os"

	"github.com/jimsmart/schema"
	_ "github.com/mattn/go-sqlite3"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("schema", func() {
	Context("with SQLite (github.com/mattn/go-sqlite3)", func() {

		const createDDL = `
			CREATE TABLE IF NOT EXISTS resource (
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
			);
			CREATE INDEX IF NOT EXISTS idx_resource_url ON resource(url);
			CREATE INDEX IF NOT EXISTS idx_resource_created_at ON resource(created_at);
			CREATE INDEX IF NOT EXISTS idx_resource_modified_at ON resource(modified_at);
			CREATE VIEW resource_view AS SELECT id, url FROM resource;
		`

		setup := func() (*sql.DB, func()) {
			path := "./test.sqlite"
			db, err := sql.Open("sqlite3", path)
			if err != nil {
				log.Fatalf("sql.Open error %v", err)
			}

			_, err = db.Exec(createDDL)
			if err != nil {
				log.Fatalf("db.Exec (create) error %v", err)
			}

			doneFn := func() {
				err = db.Close()
				if err != nil {
					log.Printf("db.Close error %v", err)
				}
				err = os.Remove(path)
				if err != nil {
					log.Printf("os.Remove error %v", err)
				}
			}

			return db, doneFn
		}

		Describe("Table", func() {
			It("should return the column type info for the table", func() {
				db, done := setup()
				defer done()
				ci, err := schema.Table(db, "resource")
				Expect(err).To(BeNil())
				Expect(ci).To(HaveLen(10))
				var list []string
				for _, c := range ci {
					list = append(list, c.Name()+" "+c.DatabaseTypeName())
				}
				Expect(list).To(Equal([]string{
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
				}))
			})
		})

		Describe("TableNames", func() {
			It("should return the table names", func() {
				db, done := setup()
				defer done()
				sn, err := schema.TableNames(db)
				Expect(err).To(BeNil())
				Expect(sn).To(HaveLen(1))
				Expect(sn).To(Equal([]string{"resource"}))
			})
		})

		Describe("Tables", func() {
			It("should return the column type info for all tables", func() {
				db, done := setup()
				defer done()
				sc, err := schema.Tables(db)
				Expect(err).To(BeNil())
				Expect(sc).To(HaveLen(1))
				ci, ok := sc["resource"]
				Expect(ok).To(BeTrue())
				Expect(ci).To(HaveLen(10))
			})
		})

		Describe("View", func() {
			It("should return the column type info for the view", func() {
				db, done := setup()
				defer done()
				ci, err := schema.View(db, "resource_view")
				Expect(err).To(BeNil())
				Expect(ci).To(HaveLen(2))
				var list []string
				for _, c := range ci {
					list = append(list, c.Name()+" "+c.DatabaseTypeName())
				}
				Expect(list).To(Equal([]string{
					"id INTEGER",
					"url TEXT",
				}))
			})
		})

		Describe("ViewNames", func() {
			It("should return the view names", func() {
				db, done := setup()
				defer done()
				sn, err := schema.ViewNames(db)
				Expect(err).To(BeNil())
				Expect(sn).To(HaveLen(1))
				Expect(sn).To(Equal([]string{"resource_view"}))
			})
		})

		Describe("Views", func() {
			It("should return the column type info for all view", func() {
				db, done := setup()
				defer done()
				sc, err := schema.Views(db)
				Expect(err).To(BeNil())
				Expect(sc).To(HaveLen(1))
				ci, ok := sc["resource_view"]
				Expect(ok).To(BeTrue())
				Expect(ci).To(HaveLen(2))
			})
		})

	})
})
