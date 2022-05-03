package store

import "github.com/belanenko/coingecko-parser/internal/app/model"

type HistoryRepository interface {
	Add(currencyName string, priceHistory []model.History) error
	GetHistory(currencyName string) ([]model.History, error)
	GetCurrencyId(currencyName string) (int, error)
}
