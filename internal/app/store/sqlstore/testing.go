package sqlstore

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/jackc/pgx/v4"
)

func TestDB(t *testing.T, databaseURL string) (*pgx.Conn, func(...string)) {
	t.Helper()

	db, err := pgx.Connect(context.Background(), databaseURL)
	if err != nil {
		t.Fatal(err)
	}

	if err := db.Ping(context.Background()); err != nil {
		t.Fatal(err)
	}

	return db, func(tables ...string) {
		defer db.Close(context.Background())
		if len(tables) > 0 {
			db.Exec(context.Background(), fmt.Sprintf("TRUNCATE %s CASCADE", strings.Join(tables, ",")))
		}
	}
}
