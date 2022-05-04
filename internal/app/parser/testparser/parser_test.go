package testparser

import (
	"testing"

	"github.com/belanenko/coingecko-parser/internal/app/model"
	"github.com/stretchr/testify/assert"
)

func TestParser_GetPriceHistoryPeriod(t *testing.T) {
	p := New()
	h, err := p.GetPriceHistoryPeriod(model.TestCurrencyName, "365")
	assert.NoError(t, err)
	assert.NotNil(t, h)
}

func TestParser_GetCurrrncyList(t *testing.T) {
	p := New()
	list := p.CurrenciesList()
	assert.Equal(t, len(list), 1)
}
