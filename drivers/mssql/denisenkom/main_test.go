package main_test

import (
	"fmt"
	"testing"

	_ "github.com/denisenkom/go-mssqldb" // DriverName: mssql

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/jimsmart/schema/drivers"
)

var _ = Describe("schema", func() {

	Context("using Microsoft SQL-Server", func() {
		// TODO(js) De-dupe this.
		const (
			user = "mssql_test_user"
			pass = "Password-123"
			host = "localhost"
			port = "41433"
		)

		Context("using driver github.com/denisenkom/go-mssqldb", func() {
			var params = drivers.MssqlDialect
			params.DriverName = "mssql"
			params.ConnStr = fmt.Sprintf("user id=%s;password=%s;server=%s;port=%s", user, pass, host, port)
			drivers.SchemaTestRunner(&params)
		})
	})
})

func TestSchema(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Schema Suite")
}
