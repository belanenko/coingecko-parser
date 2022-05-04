package parser

import "github.com/belanenko/coingecko-parser/internal/app/model"

type Parser interface {
	GetPriceHistoryPeriod(currencyName string, days string) ([]model.History, error)
	Len() int
	CurrenciesList() []string
}
