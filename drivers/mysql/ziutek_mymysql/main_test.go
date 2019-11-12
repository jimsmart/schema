package main_test

import (
	"fmt"
	"testing"

	"github.com/jimsmart/schema/drivers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	_ "github.com/ziutek/mymysql/godrv" // DriverName: mymysql
)

func TestDriver(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Driver github.com/ziutek/mymysql/godrv (mysql)")
}

var _ = Describe("Using driver github.com/ziutek/mymysql/godrv (mysql)", func() {

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
