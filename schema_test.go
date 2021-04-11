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

	TableExpRes []string
	ViewExpRes  []string

	TableNamesExpRes [][2]string
	ViewNamesExpRes  [][2]string

	PrimaryKeysExpRes []string
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
				// log.Fatalf("db.Exec (create) error %v", err)
				log.Printf("db.Exec (create) error %v exec %s", err, ddl)
			}
		}

		doneFn := func() {
			for _, ddl := range params.DropDDL {
				_, err = db.Exec(ddl)
				if err != nil {
					// log.Fatalf("db.Exec (drop) error %v", err)
					log.Printf("db.Exec (drop) error %v exec %s", err, ddl)
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
			ci, err := schema.Table(db, params.TableNamesExpRes[1][0], params.TableNamesExpRes[1][1])
			Expect(err).To(BeNil())
			var list []string
			for _, c := range ci {
				list = append(list, c.Name())
			}
			Expect(list).To(Equal(params.TableExpRes))
		})
		// TODO(js) check with empty schema param
		It("should return an error for a non-existing table", func() {
			db, done := setup()
			defer done()
			_, err := schema.Table(db, "", "XXX-NO-SUCH-TABLE-XXX")
			Expect(err).ToNot(BeNil())
		})
	})

	Describe("TableNames", func() {
		It("should return the table names", func() {
			db, done := setup()
			defer done()
			sn, err := schema.TableNames(db)
			Expect(err).To(BeNil())
			Expect(sn).To(Equal(params.TableNamesExpRes))
		})
	})

	Describe("Tables", func() {
		It("should return the column type info for all tables", func() {
			db, done := setup()
			defer done()
			sc, err := schema.Tables(db)
			Expect(err).To(BeNil())
			Expect(sc).To(HaveLen(2))
			// TODO(js) Improve / cleanup tests.
			// Expect(sc).To(HaveKey())
			ci, ok := sc[params.TableNamesExpRes[1]]
			Expect(ok).To(BeTrue())
			Expect(ci).To(HaveLen(10))
		})
	})

	Describe("View", func() {
		It("should return the column type info for the view", func() {
			db, done := setup()
			defer done()
			ci, err := schema.View(db, params.ViewNamesExpRes[0][0], params.ViewNamesExpRes[0][1])
			Expect(err).To(BeNil())
			var list []string
			for _, c := range ci {
				list = append(list, c.Name())
			}
			Expect(list).To(Equal(params.ViewExpRes))
		})
		// TODO(js) check with empty schema param
	})

	Describe("ViewNames", func() {
		It("should return the view names", func() {
			db, done := setup()
			defer done()
			sn, err := schema.ViewNames(db)
			Expect(err).To(BeNil())
			Expect(sn).To(Equal(params.ViewNamesExpRes))
		})
	})

	Describe("Views", func() {
		It("should return the column type info for all views", func() {
			db, done := setup()
			defer done()
			sc, err := schema.Views(db)
			Expect(err).To(BeNil())
			Expect(sc).To(HaveLen(1))
			ci, ok := sc[params.ViewNamesExpRes[0]]
			Expect(ok).To(BeTrue())
			Expect(ci).To(HaveLen(2))
		})
	})

	Describe("PrimaryKey", func() {
		It("should return the primary key", func() {
			db, done := setup()
			defer done()
			pk, err := schema.PrimaryKey(db, params.TableNamesExpRes[0][0], params.TableNamesExpRes[0][1])
			Expect(err).To(BeNil())
			Expect(pk).To(Equal(params.PrimaryKeysExpRes))
		})
		// TODO(js) check with empty schema param
	})

}

var _ = Describe("schema", func() {
	Context("using an unsupported (fake) db driver", func() {
		sql.Register("fakedb", FakeDb{})
		db, _ := sql.Open("fakedb", "")

		It("should return errors for every method", func() {

			var unknownDriverErr = schema.UnknownDriverError{Driver: "schema_test.FakeDb"}

			ci, err := schema.Table(db, "", "web_resource")
			Expect(ci).To(BeNil())
			Expect(err).To(MatchError(unknownDriverErr))

			tn, err := schema.TableNames(db)
			Expect(tn).To(BeNil())
			Expect(err).To(MatchError(unknownDriverErr))

			// ta, err := schema.Tables(db)
			// Expect(ta).To(BeNil())
			// Expect(err).To(MatchError(unknownDriverErr))

			ci, err = schema.View(db, "", "web_resource")
			Expect(ci).To(BeNil())
			Expect(err).To(MatchError(unknownDriverErr))

			vn, err := schema.ViewNames(db)
			Expect(vn).To(BeNil())
			Expect(err).To(MatchError(unknownDriverErr))

			// vw, err := schema.Views(db)
			// Expect(vw).To(BeNil())
			// Expect(err).To(MatchError(unknownDriverErr))
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
