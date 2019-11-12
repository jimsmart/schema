package main_test

import (
	"fmt"
	"testing"

	"github.com/jimsmart/schema/drivers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	_ "gopkg.in/goracle.v2" // DriverName: goracle
)

func TestDriver(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Driver github.com/go-goracle/goracle (oracle)")
}

var _ = Describe("Using driver github.com/go-goracle/goracle (oracle)", func() {

	// TODO(js) De-dupe this.
	const (
		user = "oracle_test_user"
		pass = "password"
		host = "localhost"
		port = "41521"
		dbs  = "xe"
	)

	var params = drivers.OracleDialect
	params.DriverName = "goracle"
	params.ConnStr = fmt.Sprintf("%s/%s@%s:%s/%s", user, pass, host, port, dbs)

	drivers.SchemaTestRunner(&params)
})
