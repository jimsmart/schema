package main_test

import (
	"fmt"
	"testing"

	"github.com/jimsmart/schema/drivers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	_ "github.com/go-sql-driver/mysql" // DriverName: mysql
)

func TestDriver(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Driver github.com/go-sql-driver/mysql (mysql)")
}

var _ = Describe("Using driver github.com/go-sql-driver/mysql (mysql)", func() {

	// TODO(js) De-dupe this.
	const (
		user = "mysql_test_user"
		pass = "password"
		host = "localhost"
		port = "43306"
		dbs  = "test_db"
	)

	var params = drivers.MysqlDialect
	params.DriverName = "mysql"
	params.ConnStr = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, pass, host, port, dbs)
	drivers.SchemaTestRunner(&params)
})
