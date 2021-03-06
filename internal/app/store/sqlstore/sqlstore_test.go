package sqlstore_test

import (
	"os"
	"testing"
)

var (
	databaseURL string
)

func TestMain(m *testing.M) {
	if os.Getenv("TEST_DATABASE_URL") == "" {
		databaseURL = "postgres://postgres:password@localhost:5433/history_test?sslmode=disable"
	}
	os.Exit(m.Run())
}
