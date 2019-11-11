package main_test

import (
	"fmt"

	"github.com/jimsmart/schema/drivers"
	. "github.com/onsi/ginkgo"

	_ "github.com/denisenkom/go-mssqldb" // DriverName: mssql
)

var _ = Describe("driver github.com/denisenkom/go-mssqldb", func() {

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
