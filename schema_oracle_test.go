package schema_test

import (
	"database/sql"
	"log"

	"github.com/jimsmart/schema"
	_ "gopkg.in/goracle.v2"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("schema", func() {
	Context("with Oracle (github.com/go-goracle/goracle)", func() {

		if true {

			// Oracle and MS-SQL both barf on multiple-batch statements.

			createDDL := []string{`
				CREATE TABLE web_resource (
					id				NUMBER NOT NULL,
					url				NVARCHAR2(1024) NOT NULL UNIQUE,
					content			BLOB,
					compressed_size	NUMBER NOT NULL,
					content_length	NUMBER NOT NULL,
					content_type	NVARCHAR2(128) NOT NULL,
					etag			NVARCHAR2(128) NOT NULL,
					last_modified	NVARCHAR2(128) NOT NULL,
					created_at		TIMESTAMP WITH TIME ZONE NOT NULL,
					modified_at		TIMESTAMP WITH TIME ZONE,
					PRIMARY KEY (id)
				)`,
				// `CREATE INDEX idx_web_resource_url ON web_resource(url)`,
				`CREATE INDEX idx_web_resource_created_at ON web_resource(created_at)`,
				`CREATE INDEX idx_web_resource_modified_at ON web_resource(modified_at)`,
				`CREATE VIEW web_resource_view AS SELECT id, url FROM web_resource`,
			}

			dropDDL := []string{
				`DROP VIEW web_resource_view`,
				`DROP INDEX idx_web_resource_modified_at`,
				`DROP INDEX idx_web_resource_created_at`,
				// `DROP INDEX idx_web_resource_url`,
				`DROP TABLE web_resource`,
			}

			setup := func() (*sql.DB, func()) {
				connStr := "test_user/password@localhost:32786/xe" // docker
				db, err := sql.Open("goracle", connStr)
				if err != nil {
					log.Fatalf("sql.Open error %v", err)
				}

				for _, ddl := range createDDL {
					_, err = db.Exec(ddl)
					if err != nil {
						// log.Fatalf("db.Exec (create) error %v", err)
						log.Printf("db.Exec (create) error %v", err)
					}
				}

				doneFn := func() {
					_ = dropDDL
					for _, ddl := range dropDDL {
						_, err = db.Exec(ddl)
						if err != nil {
							// log.Fatalf("db.Exec (drop) error %v", err)
							log.Printf("db.Exec (drop) error %v", err)
						}
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
					ci, err := schema.Table(db, "web_resource")
					Expect(err).To(BeNil())
					Expect(ci).To(HaveLen(10))
					var list []string
					for _, c := range ci {
						list = append(list, c.Name()+" "+c.DatabaseTypeName())
					}
					Expect(list).To(Equal([]string{
						"ID NUMBER",
						"URL NVARCHAR2",
						"CONTENT BLOB",
						"COMPRESSED_SIZE NUMBER",
						"CONTENT_LENGTH NUMBER",
						"CONTENT_TYPE NVARCHAR2",
						"ETAG NVARCHAR2",
						"LAST_MODIFIED NVARCHAR2",
						"CREATED_AT TIMESTAMP WITH TIMEZONE",
						"MODIFIED_AT TIMESTAMP WITH TIMEZONE",
					}))
				})
			})

			Describe("TableNames", func() {
				It("should return the table names", func() {
					db, done := setup()
					defer done()

					// err := oraDump(db)
					// Expect(err).To(BeNil())

					sn, err := schema.TableNames(db)
					Expect(err).To(BeNil())
					Expect(sn).To(HaveLen(1))
					Expect(sn).To(Equal([]string{"WEB_RESOURCE"}))
				})
			})

			Describe("Tables", func() {
				It("should return the column type info for all tables", func() {
					db, done := setup()
					defer done()
					sc, err := schema.Tables(db)
					Expect(err).To(BeNil())
					Expect(sc).To(HaveLen(1))
					ci, ok := sc["WEB_RESOURCE"]
					Expect(ok).To(BeTrue())
					Expect(ci).To(HaveLen(10))
				})
			})

			Describe("View", func() {
				It("should return the column type info for the view", func() {
					db, done := setup()
					defer done()
					ci, err := schema.View(db, "web_resource_view")
					Expect(err).To(BeNil())
					Expect(ci).To(HaveLen(2))
					var list []string
					for _, c := range ci {
						list = append(list, c.Name()+" "+c.DatabaseTypeName())
					}
					Expect(list).To(Equal([]string{
						"ID NUMBER",
						"URL NVARCHAR2",
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
					Expect(sn).To(Equal([]string{"WEB_RESOURCE_VIEW"}))
				})
			})

			Describe("Views", func() {
				It("should return the column type info for all views", func() {
					db, done := setup()
					defer done()
					sc, err := schema.Views(db)
					Expect(err).To(BeNil())
					Expect(sc).To(HaveLen(1))
					ci, ok := sc["WEB_RESOURCE_VIEW"]
					Expect(ok).To(BeTrue())
					Expect(ci).To(HaveLen(2))
				})
			})

		}

	})
})

// func oraDump(db *sql.DB) error {

// 	//SELECT table_name FROM user_tables
// 	rows, err := db.Query(`
// 		SELECT *
// 		  FROM user_tables
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
// 			// return err
// 			log.Printf("%v", err)
// 		}
// 		s := ""
// 		for _, v := range vals {
// 			s = s + fmt.Sprintf("%s ", v)
// 		}
// 		log.Print(s)
// 	}
// 	return nil
// }
