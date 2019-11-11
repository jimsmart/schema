package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestMinus5(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Driver github.com/minus5/gofreetds (mssql)")
}
