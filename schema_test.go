package schema_test

import (
	"database/sql"
	"database/sql/driver"
	"log"
	"strings"

	"github.com/jimsmart/schema"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type testParams struct {
	DriverName string
	ConnStr    string

	CreateDDL []string
	DropDDL   []string
	DropFn    func()

	TableNamesExpRes []string
	ViewNamesExpRes  []string
}

func SchemaTestRunner(params *testParams) {

	setup := func() (*sql.DB, func()) {
		db, err := sql.Open(params.DriverName, params.ConnStr)
		if err != nil {
			log.Fatalf("sql.Open error %v", err)
		}

		for _, ddl := range params.CreateDDL {
			_, err = db.Exec(ddl)
			if err != nil {
				log.Printf("db.Exec (create) error %v exec %s", err, ddl)
			}
		}

		doneFn := func() {
			for _, ddl := range params.DropDDL {
				_, err = db.Exec(ddl)
				if err != nil {
					log.Printf("db.Exec (drop) error %v exec %s", err, ddl)
				}
			}

			// Ensure our tests have properly deleted their tables and views.
			tableNames, err := schema.TableNames(db)
			if err != nil {
				log.Printf("schema.TableNames (after drop) error %v", err)
			}
			if len(tableNames) != 0 {
				log.Println("schema.TableNames reports undeleted tables (test has bad drop ddl?)")
				for _, name := range tableNames {
					log.Println(name)
				}
			}
			viewNames, err := schema.ViewNames(db)
			if err != nil {
				log.Printf("schema.ViewNames (after drop) error %v", err)
			}
			if len(viewNames) != 0 {
				log.Println("schema.ViewNames reports undeleted views (test has bad drop ddl?)")
				for _, name := range viewNames {
					log.Println(name)
				}
			}

			err = db.Close()
			if err != nil {
				log.Printf("db.Close error %v", err)
			}
			if params.DropFn != nil {
				params.DropFn()
			}
		}

		return db, doneFn
	}

	Describe("Table", func() {
		It("should return the column type info for an existing table", func() {
			db, done := setup()
			defer done()
			ct, err := schema.Table(db, "web_resource")
			Expect(err).To(BeNil())
			Expect(ct).To(HaveLen(10))
		})
		It("should return an error for a non-existing table", func() {
			db, done := setup()
			defer done()
			_, err := schema.Table(db, "XXX-NO-SUCH-TABLE-XXX")
			Expect(err).ToNot(BeNil())
		})
		It("should handle tables with unusual names (escaping)", func() {
			tableNames := []string{
				"blanks in name",
				"[brackets] in name",
				`"d.quotes" in name`,
				"'s.quotes' in name",
				"{braces} in name",
				"`backticks` in name",
				`backslashes\in\name`,
			}
			db, done := setup()
			defer done()
			for _, tn := range tableNames {
				ct, err := schema.Table(db, tn)
				Expect(err).To(BeNil())
				Expect(ct).To(HaveLen(1))
			}
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
			Expect(sn).To(ConsistOf(params.TableNamesExpRes))
		})

		It("should return no table names for an empty database", func() {
			db, err := sql.Open(params.DriverName, params.ConnStr)
			Expect(err).To(BeNil())
			defer func() {
				err = db.Close()
				Expect(err).To(BeNil())
				if params.DropFn != nil {
					params.DropFn()
				}
			}()
			sn, err := schema.TableNames(db)
			Expect(err).To(BeNil())
			Expect(sn).To(HaveLen(0))
		})
	})

	Describe("Tables", func() {
		It("should return the column type info for all tables", func() {
			db, done := setup()
			defer done()
			sc, err := schema.Tables(db)
			Expect(err).To(BeNil())
			Expect(sc).To(HaveLen(len(params.TableNamesExpRes)))
			ct, ok := sc["web_resource"]
			Expect(ok).To(BeTrue())
			Expect(ct).To(HaveLen(10))
		})
	})

	Describe("View", func() {
		It("should return the column type info for the view", func() {
			db, done := setup()
			defer done()
			ct, err := schema.View(db, "web_resource_view")
			Expect(err).To(BeNil())
			Expect(ct).To(HaveLen(2))
		})
	})

	Describe("ViewNames", func() {
		It("should return the view names", func() {
			db, done := setup()
			defer done()
			sn, err := schema.ViewNames(db)
			Expect(err).To(BeNil())
			Expect(sn).To(ConsistOf(params.ViewNamesExpRes))
		})

		It("should return no view names for an empty database", func() {
			db, err := sql.Open(params.DriverName, params.ConnStr)
			Expect(err).To(BeNil())
			defer func() {
				err = db.Close()
				Expect(err).To(BeNil())
				if params.DropFn != nil {
					params.DropFn()
				}
			}()
			sn, err := schema.ViewNames(db)
			Expect(err).To(BeNil())
			Expect(sn).To(HaveLen(0))
		})
	})

	Describe("Views", func() {
		It("should return the column type info for all views", func() {
			db, done := setup()
			defer done()
			sc, err := schema.Views(db)
			Expect(err).To(BeNil())
			Expect(sc).To(HaveLen(len(params.ViewNamesExpRes)))
			ct, ok := sc["web_resource_view"]
			Expect(ok).To(BeTrue())
			Expect(ct).To(HaveLen(2))
		})
	})

}

var _ = Describe("schema", func() {
	Context("using an unsupported (fake) db driver", func() {
		sql.Register("fakedb", FakeDb{})
		db, _ := sql.Open("fakedb", "")

		It("should return errors for every method", func() {

			var unknownDriverErr = schema.UnknownDriverError{Driver: "schema_test.FakeDb"}

			ct, err := schema.Table(db, "web_resource")
			Expect(ct).To(BeNil())
			Expect(err).To(MatchError(unknownDriverErr))

			tn, err := schema.TableNames(db)
			Expect(tn).To(BeNil())
			Expect(err).To(MatchError(unknownDriverErr))

			ta, err := schema.Tables(db)
			Expect(ta).To(BeNil())
			Expect(err).To(MatchError(unknownDriverErr))

			ct, err = schema.View(db, "web_resource")
			Expect(ct).To(BeNil())
			Expect(err).To(MatchError(unknownDriverErr))

			vn, err := schema.ViewNames(db)
			Expect(vn).To(BeNil())
			Expect(err).To(MatchError(unknownDriverErr))

			vw, err := schema.Views(db)
			Expect(vw).To(BeNil())
			Expect(err).To(MatchError(unknownDriverErr))

			errm := err.Error()
			Expect(errm).To(ContainSubstring("unknown"))
			Expect(errm).To(ContainSubstring(" schema_test.FakeDb"))
		})
	})
})

type FakeDb struct{}

func (_ FakeDb) Open(name string) (driver.Conn, error) {
	return nil, nil
}

// pack a string, normalising its whitespace.
func pack(s string) string {
	return strings.Join(strings.Fields(s), " ")
}
