package teststore

import (
	"github.com/belanenko/coingecko-parser/internal/app/model"
	"github.com/belanenko/coingecko-parser/internal/app/store"
)

type HistoryRepository struct {
	Store *Store
}

func (r *HistoryRepository) Add(currencyName string, priceHistory []model.History) error {
	r.Store.db[currencyName] = priceHistory
	return nil
}

func (r *HistoryRepository) GetHistory(currencyName string) ([]model.History, error) {
	return r.Store.db[currencyName], nil
}

func (r *HistoryRepository) GetCurrencyId(currencyName string) (int, error) {
	_, ok := r.Store.db[currencyName]
	if !ok {
		return -1, store.ErrNoRows
	}
	return 0, nil
}
