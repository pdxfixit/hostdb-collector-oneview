package main

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {

	loadConfig()

	os.Exit(m.Run())

}
