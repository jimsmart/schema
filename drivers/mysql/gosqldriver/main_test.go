package main_test

import (
	"fmt"

	"github.com/jimsmart/schema/drivers"
	. "github.com/onsi/ginkgo"

	_ "github.com/go-sql-driver/mysql" // DriverName: mysql
)

var _ = Describe("driver github.com/go-sql-driver/mysql", func() {

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
