package sqlstore_test

import (
	"context"
	"testing"

	"github.com/belanenko/coingecko-parser/internal/app/store/sqlstore"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestMigrator_UP_DOWN(t *testing.T) {
	m := sqlstore.NewMigrator(databaseURL)
	assert.NoError(t, m.Up())
	store, _ := sqlstore.TestDB(t, databaseURL)

	_, err := store.Exec(context.Background(), "SELECT * FROM price_history;")
	assert.NoError(t, err)
	assert.NoError(t, m.Down())

}
