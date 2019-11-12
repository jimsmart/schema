package main_test

import (
	"fmt"
	"testing"

	"github.com/jimsmart/schema/drivers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	_ "github.com/denisenkom/go-mssqldb" // DriverName: mssql
)

func TestDriver(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Driver github.com/denisenkom/go-mssqldb (mssql)")
}

var _ = Describe("Using driver github.com/denisenkom/go-mssqldb (mssql)", func() {

	// TODO(js) De-dupe this.
	const (
		user = "mssql_test_user"
		pass = "Password-123"
		host = "localhost"
		port = "41433"
	)

	var params = drivers.MssqlDialect
	params.DriverName = "mssql"
	params.ConnStr = fmt.Sprintf("user id=%s;password=%s;server=%s;port=%s", user, pass, host, port)
	drivers.SchemaTestRunner(&params)
})
