package main

import (
	"authentication/data"
	"os"
	"testing"
)

var testApp Config

func TestMain(m *testing.M) {
	// exit after running all tests
	repo := data.NewPostgresTestRepository(nil)
	testApp.Repo = repo
	os.Exit(m.Run())
}
