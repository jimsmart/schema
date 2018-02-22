package schema_test

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jimsmart/schema"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("schema", func() {
	Context("with MySQL (github.com/go-sql-driver/mysql)", func() {

		if true {

			const createDDL = `
			CREATE TABLE IF NOT EXISTS resource (
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
			);
			CREATE VIEW resource_view AS SELECT id, url FROM resource;
		`
			const dropDDL = `
			DROP VIEW resource_view;
			DROP TABLE resource;
		`

			setup := func() (*sql.DB, func()) {
				// username:password@protocol(address)/dbname?param=value
				connStr := "test_user:password@tcp(localhost:32778)/test_db?multiStatements=true" // docker
				db, err := sql.Open("mysql", connStr)
				if err != nil {
					log.Fatalf("sql.Open error %v", err)
				}

				_, err = db.Exec(createDDL)
				if err != nil {
					log.Fatalf("db.Exec (create) error %v", err)
				}

				doneFn := func() {
					_, err = db.Exec(dropDDL)
					if err != nil {
						log.Printf("db.Exec (drop) error %v", err)
					}
					err = db.Close()
					if err != nil {
						log.Printf("db.Close error %v", err)
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
						"id INT",
						"url VARCHAR",
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
				It("should return the column type info for all views", func() {
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

		}

	})
})
