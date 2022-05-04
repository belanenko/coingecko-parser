package model

import "testing"

var (
	TestCurrencyName = "testcoin"
)

func TestCurrency(t *testing.T) *Currencies {
	t.Helper()

	currency := &Currencies{
		Items: map[string][]History{
			"testcoin": {
				{
					Timestamp: "1651603202",
					Price:     4500.23,
				},
			},
		},
	}

	return currency

}
