package main_test

import (
	"fmt"
	"testing"

	_ "github.com/minus5/gofreetds" // DriverName: mssql

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/jimsmart/schema/drivers"
)

var _ = XDescribe("schema", func() {

	Context("using Microsoft SQL-Server", func() {
		// TODO(js) De-dupe this.
		const (
			user = "mssql_test_user"
			pass = "Password-123"
			host = "localhost"
			port = "41433"
		)

		// TODO(js) This driver (gofreetds) seems buggy. :/
		Context("using driver github.com/minus5/gofreetds", func() {

			var params = drivers.MssqlDialect
			params.DriverName = "mssql"
			params.ConnStr = fmt.Sprintf("user id=%s;password=%s;server=%s:%s", user, pass, host, port)
			drivers.SchemaTestRunner(&params)
		})
	})
})

func TestSchema(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Schema Suite")
}
