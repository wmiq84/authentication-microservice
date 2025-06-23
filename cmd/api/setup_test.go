package main

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// exit after running all tests
	os.Exit(m.Run())
}
