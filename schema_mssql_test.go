package schema_test

import (
	"database/sql"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/jimsmart/schema"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("schema", func() {
	Context("with Microsoft SQL-Server (github.com/denisenkom/go-mssqldb)", func() {

		if true {

			const createDDL = `
			CREATE TABLE resource (
				id				INTEGER NOT NULL,
				url				NVARCHAR NOT NULL UNIQUE,
				content			VARBINARY,
				compressed_size	INTEGER NOT NULL,
				content_length	INTEGER NOT NULL,
				content_type	NVARCHAR NOT NULL,
				etag			NVARCHAR NOT NULL,
				last_modified	NVARCHAR NOT NULL,
				created_at		DATETIME NOT NULL,
				modified_at		DATETIME,
				PRIMARY KEY (id)
			);
			CREATE INDEX idx_resource_url ON resource(url);
			CREATE INDEX idx_resource_created_at ON resource (created_at);
			CREATE INDEX idx_resource_modified_at ON resource (modified_at);
		`
			// TODO Hmmm this driver doesn't seem to support 'go' commands for multi-batch?
			const createDDL2 = `
			CREATE VIEW resource_view AS SELECT id, url FROM resource;
		`

			const dropDDL = `
			DROP VIEW resource_view;
			DROP INDEX IF EXISTS idx_resource_modified_at ON resource;
			DROP INDEX IF EXISTS idx_resource_created_at ON resource;
			DROP INDEX IF EXISTS idx_resource_url ON resource;
			DROP TABLE resource;
		`

			setup := func() (*sql.DB, func()) {
				connStr := "server=localhost;port=32784;user id=test_user;password=aNRV!^5-WCe4hz$3" // docker
				db, err := sql.Open("mssql", connStr)
				if err != nil {
					log.Fatalf("sql.Open error %v", err)
				}

				_, err = db.Exec(createDDL)
				if err != nil {
					log.Fatalf("db.Exec (create) error %v", err)
				}

				_, err = db.Exec(createDDL2)
				if err != nil {
					log.Fatalf("db.Exec (create2) error %v", err)
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
						"url NVARCHAR",
						"content VARBINARY",
						"compressed_size INT",
						"content_length INT",
						"content_type NVARCHAR",
						"etag NVARCHAR",
						"last_modified NVARCHAR",
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
				It("should return the  column type info for all tables", func() {
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
						"url NVARCHAR",
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
				It("should return the  column type info for all views", func() {
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

// func msDump(db *sql.DB) error {

// 	rows, err := db.Query(`
// 		SELECT S.name as Owner, T.name as TableName
// 		FROM
// 		  sys.tables AS T
// 		    INNER JOIN sys.schemas AS S ON S.schema_id = T.schema_id
// 		    LEFT JOIN sys.extended_properties AS EP ON EP.major_id = T.[object_id]
// 		WHERE
// 		  T.is_ms_shipped = 0 AND
// 		  (EP.class_desc IS NULL OR (EP.class_desc <>'OBJECT_OR_COLUMN' AND
// 		  EP.[name] <> 'microsoft_database_tools_support'))
// 	  `)

// 	// rows, err := db.Query(`
// 	// 	SELECT *
// 	// 	  FROM information_schema.tables
// 	// 	 WHERE table_type='BASE TABLE'
// 	// 	   -- AND table_schema='public'
// 	//   `)

// 	// rows, err := db.Query(`
// 	// 	SELECT name, type
// 	// 	  FROM sysobjects
// 	// 	 WHERE xtype = 'U'
// 	// 	  -- table_type='BASE TABLE'
// 	// 	   -- AND table_schema='public'
// 	//   `)
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
