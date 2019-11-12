package main_test

import (
	"testing"

	"github.com/jimsmart/schema/drivers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	_ "github.com/mattn/go-sqlite3" // DriverName: sqlite3
)

func TestDriver(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Driver github.com/mattn/go-sqlite3 (sqlite)")
}

var _ = Describe("Using driver github.com/mattn/go-sqlite3 (sqlite)", func() {

	// TODO(js) De-dupe this.
	const (
		// user = ""
		// pass = ""
		// host = ""
		// port = ""
		dbs = ":memory:"
		// dbs = "./test.sqlite"
	)

	var params = drivers.SqliteDialect
	params.DriverName = "sqlite3"
	params.ConnStr = dbs

	drivers.SchemaTestRunner(&params)
})
