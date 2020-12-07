package database

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	Setup()
	result := m.Run()
	TearDown()
	os.Exit(result)
}
