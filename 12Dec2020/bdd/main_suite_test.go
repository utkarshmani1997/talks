package main

import (
	"flag"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	testEnvironment string
)

func init() {
	flag.StringVar(&testEnvironment, "testEnv", "local", "testing environment local/pipeline")
}

func TestGinkgo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Ginkgo Suite")
}

var _ = BeforeSuite(func() {})
var _ = AfterSuite(func() {})
