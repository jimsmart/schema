package main_test

import (
	"fmt"

	"github.com/jimsmart/schema/drivers"
	. "github.com/onsi/ginkgo"

	_ "github.com/ziutek/mymysql/godrv" // DriverName: mymysql
)

var _ = Describe("driver github.com/ziutek/mymysql/godrv", func() {

	// TODO(js) De-dupe this.
	const (
		user = "mysql_test_user"
		pass = "password"
		host = "localhost"
		port = "43306"
		dbs  = "test_db"
	)

	var params = drivers.MysqlDialect
	params.DriverName = "mymysql"
	params.ConnStr = fmt.Sprintf("tcp:%s:%s*%s/%s/%s", host, port, dbs, user, pass)
	drivers.SchemaTestRunner(&params)
})
