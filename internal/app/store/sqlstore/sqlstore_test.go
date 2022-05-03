package sqlstore_test

import (
	"os"
	"testing"
)

var (
	databaseURL string
)

func TestMain(m *testing.M) {
	if os.Getenv("DATABASE_URL") == "" {
		databaseURL = "postgres://postgres:password@localhost:5432/history_test?sslmode=disable"
	}
	os.Exit(m.Run())
}
