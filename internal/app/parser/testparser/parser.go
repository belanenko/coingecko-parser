package testparser

import (
	"github.com/belanenko/coingecko-parser/internal/app/model"
)

type TestParser struct {
}

func New() *TestParser {
	return &TestParser{}
}

func (p *TestParser) GetPriceHistoryPeriod(currencyName string, days string) ([]model.History, error) {
	history := []model.History{
		{
			Timestamp: "1651603202",
			Price:     4500.23,
		},
	}

	return history, nil
}

func (p *TestParser) Len() int {
	return 1
}

func (p *TestParser) CurrenciesList() []string {
	return []string{model.TestCurrencyName}
}
