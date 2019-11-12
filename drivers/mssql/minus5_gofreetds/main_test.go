package main_test

import (
	"fmt"
	"testing"

	"github.com/jimsmart/schema/drivers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	_ "github.com/minus5/gofreetds" // DriverName: mssql
)

func TestDriver(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Driver github.com/minus5/gofreetds (mssql)")
}

var _ = XDescribe("Using driver github.com/minus5/gofreetds (mssql)", func() {

	// TODO(js) De-dupe this.
	const (
		user = "mssql_test_user"
		pass = "Password-123"
		host = "localhost"
		port = "41433"
	)

	// TODO(js) This driver (gofreetds) seems buggy. :/

	// Possibly related? https://stackoverflow.com/questions/48795459/using-tbl-with-in-schema-creating-syntax-error-using-freetds

	var params = drivers.MssqlDialect
	params.DriverName = "mssql"
	params.ConnStr = fmt.Sprintf("user id=%s;password=%s;server=%s:%s", user, pass, host, port)
	drivers.SchemaTestRunner(&params)
})
