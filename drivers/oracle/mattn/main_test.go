package main_test

import (
	"fmt"
	"testing"

	"github.com/jimsmart/schema/drivers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	_ "github.com/mattn/go-oci8" // DriverName: oci8
)

func TestDriver(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Driver github.com/mattn/go-oci8 (oracle)")
}

var _ = Describe("Using driver github.com/mattn/go-oci8 (oracle)", func() {

	// TODO(js) De-dupe this.
	const (
		user = "oracle_test_user"
		pass = "password"
		host = "localhost"
		port = "41521"
		dbs  = "xe"
	)

	var params = drivers.OracleDialect
	params.DriverName = "oci8"
	params.ConnStr = fmt.Sprintf("%s/%s@%s:%s/%s", user, pass, host, port, dbs)

	drivers.SchemaTestRunner(&params)
})
