package schema_test

import (
	"database/sql"
	"log"

	"github.com/jimsmart/schema"
	_ "github.com/lib/pq"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("schema", func() {
	Context("with Postgres (github.com/lib/pq)", func() {

		if true {

			const createDDL = `
			CREATE TABLE IF NOT EXISTS resource (
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
			);
			CREATE INDEX IF NOT EXISTS idx_resource_url ON resource(url);
			CREATE INDEX IF NOT EXISTS idx_resource_created_at ON resource(created_at);
			CREATE INDEX IF NOT EXISTS idx_resource_modified_at ON resource(modified_at);
			CREATE VIEW resource_view AS SELECT id, url FROM resource;
		`

			const dropDDL = `
			DROP VIEW IF EXISTS resource_view;
			DROP INDEX IF EXISTS idx_resource_modified_at;
			DROP INDEX IF EXISTS idx_resource_created_at;
			DROP INDEX IF EXISTS idx_resource_url;
			DROP TABLE IF EXISTS resource;
		`

			setup := func() (*sql.DB, func()) {
				connStr := "host=localhost port=32779 user=postgres dbname=postgres sslmode=disable" // docker
				db, err := sql.Open("postgres", connStr)
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
				It("should return the column type info", func() {
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
						"id INT4",
						"url TEXT",
						"content BYTEA",
						"compressed_size INT4",
						"content_length INT4",
						"content_type TEXT",
						"etag TEXT",
						"last_modified TEXT",
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
						"id INT4",
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

// func pgDump(db *sql.DB) error {

// 	rows, err := db.Query(`
// 		SELECT *
// 		  FROM information_schema.tables
// 		 WHERE table_schema='public'
// 		   AND table_type='BASE TABLE'
//    `)
// 	if err != nil {
// 		return err
// 	}
// 	defer rows.Close()

// 	ci, err := rows.ColumnTypes()
// 	if err != nil {
// 		return err
// 	}
// 	for _, c := range ci {
// 		log.Printf("%v", c)
// 	}

// 	cols, err := rows.Columns()
// 	if err != nil {
// 		return err
// 	}
// 	vals := make([]interface{}, len(cols))
// 	for i, _ := range cols {
// 		vals[i] = new(sql.RawBytes)
// 	}

// 	for rows.Next() {
// 		err = rows.Scan(vals...)
// 		if err != nil {
// 			return err
// 		}
// 		s := ""
// 		for _, v := range vals {
// 			s = s + fmt.Sprintf("%s ", v)
// 		}
// 		log.Print(s)
// 	}
// 	return nil
// }
