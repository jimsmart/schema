package main_test

import (
	"fmt"
	"testing"

	"github.com/jimsmart/schema/drivers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	_ "github.com/lib/pq" // DriverName: postgres
)

func TestDriver(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Driver github.com/lib/pq (postgres)")
}

var _ = Describe("Using driver github.com/lib/pq (postgres)", func() {

	// TODO(js) De-dupe this.
	const (
		user = "postgres"
		host = "localhost"
		port = "45432"
	)

	var params = drivers.PostgresDialect
	params.DriverName = "postgres"
	params.ConnStr = fmt.Sprintf("user=%s host=%s port=%s sslmode=disable", user, host, port)
	drivers.SchemaTestRunner(&params)
})
