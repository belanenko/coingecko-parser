package sqlstore_test

import (
	"testing"

	"github.com/belanenko/coingecko-parser/internal/app/model"
	"github.com/belanenko/coingecko-parser/internal/app/store/sqlstore"
	"github.com/stretchr/testify/assert"
)

func TestHistoryRepository_Add(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("price_history", "currencies")
	s := sqlstore.New(db)
	c := model.TestCurrency(t)
	currencyName := "testcoin"
	assert.NoError(t, s.History().Add(currencyName, c.Items[currencyName]))
}

func TestHistoryRepository_GetHistory(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("price_history", "currencies")
	s := sqlstore.New(db)
	c := model.TestCurrency(t)
	currencyName := "testcoin"
	s.History().Add("testcoin", c.Items[currencyName])

	h, err := s.History().GetHistory(currencyName)
	assert.NoError(t, err)
	assert.Equal(t, h, c.Items[currencyName])
}

func TestHistoryRepository_GetCurrencyId(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("price_history", "currencies")
	s := sqlstore.New(db)
	currencyName := "testcoin"

	id, err := s.History().GetCurrencyId(currencyName)
	assert.NoError(t, err)
	assert.Equal(t, id, -1)

	c := model.TestCurrency(t)
	s.History().Add("testcoin", c.Items[currencyName])
	id, err = s.History().GetCurrencyId(currencyName)
	assert.NoError(t, err)
	assert.NotEqual(t, id, -1)

}
