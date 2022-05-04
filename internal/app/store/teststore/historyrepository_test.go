package teststore_test

import (
	"testing"

	"github.com/belanenko/coingecko-parser/internal/app/model"
	"github.com/belanenko/coingecko-parser/internal/app/store"
	"github.com/belanenko/coingecko-parser/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
)

func TestHistoryRepository_Add(t *testing.T) {
	s := teststore.New()
	m := model.TestCurrency(t)
	assert.NoError(t, s.History().Add(model.TestCurrencyName, m.Items[model.TestCurrencyName]))
}

func TestHistoryRepository_GetHistory(t *testing.T) {
	s := teststore.New()
	m := model.TestCurrency(t)
	assert.NoError(t, s.History().Add(model.TestCurrencyName, m.Items[model.TestCurrencyName]))

	h, _ := s.History().GetHistory(model.TestCurrencyName)
	assert.Equal(t, m.Items[model.TestCurrencyName], h)
}

func TestHistoryRepository_GetCurrencyId(t *testing.T) {
	s := teststore.New()
	m := model.TestCurrency(t)

	id, err := s.History().GetCurrencyId(model.TestCurrencyName)
	assert.EqualError(t, store.ErrNoRows, err.Error())
	assert.Equal(t, id, -1)

	assert.NoError(t, s.History().Add(model.TestCurrencyName, m.Items[model.TestCurrencyName]))
	id, err = s.History().GetCurrencyId(model.TestCurrencyName)
	assert.Nil(t, err)
	assert.NotEqual(t, id, -1)
}
