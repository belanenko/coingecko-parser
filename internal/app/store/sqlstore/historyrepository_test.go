package sqlstore_test

import (
	"testing"

	"github.com/belanenko/coingecko-parser/internal/app/model"
	"github.com/belanenko/coingecko-parser/internal/app/store"
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

// TODO: переписать тест
func TestHistoryRepository_GetCurrencyId(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("price_history", "currencies")
	currencyName := "testcoin"
	s := sqlstore.New(db)
	c := model.TestCurrency(t)
	assert.NoError(t, s.History().Add("testcoin", c.Items[currencyName]))

	testCases := []struct {
		name         string
		currencyName string
		valid        bool
	}{
		{
			name:         "valid",
			currencyName: currencyName,
			valid:        true,
		},
		{
			name:         "invalid",
			currencyName: "invalid",
			valid:        false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			id, err := s.History().GetCurrencyId(tc.currencyName)
			if err != nil {
				if tc.valid {
					t.Fatal(err)
				} else {
					assert.EqualError(t, store.ErrNoRows, err.Error())
				}
			}
			if tc.valid {
				assert.NotEqual(t, -1, id)
			} else {
				assert.Equal(t, -1, id)
			}
		})
	}
}
