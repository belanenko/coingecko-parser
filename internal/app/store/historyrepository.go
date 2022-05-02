package store

import (
	"context"
	"database/sql"
	"errors"
	"strconv"

	"github.com/belanenko/coingecko-parser/internal/app/model"
)

type HistoryRepository struct {
	Store *Store
}

func (h *HistoryRepository) Add(currencyName string, priceHistory []model.History) error {
	currencyId, err := h.GetCurrencyId(currencyName)
	if err != nil {
		if err != sql.ErrNoRows {
			return err
		}
	}
	if currencyId == -1 {
		if err := h.Store.db.QueryRow(context.Background(), "INSERT INTO currencies (name) VALUES ($1) RETURNING id;", currencyName).Scan(&currencyId); err != nil {
			return err
		}
	}

	for _, history := range priceHistory {
		if _, err := h.Store.db.Exec(
			context.Background(), "INSERT INTO price_history (tstamp, price, fk_currencies) VALUES ($1, $2, $3);",
			history.Timestamp, history.Price, currencyId,
		); err != nil {
			return errors.New("error inserting history")
		}
	}
	return nil
}

func (h *HistoryRepository) GetHistory(currencyName string) ([]model.History, error) {
	currencyId, err := h.GetCurrencyId(currencyName)
	if err != nil {
		return nil, err
	}

	rows, err := h.Store.db.Query(context.Background(), "SELECT tstamp, price FROM price_history WHERE fk_currencies=$1;", currencyId)
	if err != nil {
		return nil, err
	}

	history := make([]model.History, 0, 365)
	for rows.Next() {
		var h model.History
		var tstamp int
		if err := rows.Scan(&tstamp, &h.Price); err != nil {
			return history, err
		}
		h.Timestamp = strconv.Itoa(tstamp)
		history = append(history, h)
	}

	return history, nil
}

func (h *HistoryRepository) GetCurrencyId(currencyName string) (int, error) {
	var currencyId int

	if err := h.Store.db.QueryRow(context.Background(), "SELECT id FROM currencies WHERE currencies.name=$1;", currencyName).Scan(&currencyId); err != nil {
		if err.Error() != "no rows in result set" {
			return -1, err
		}
		return -1, nil
	}
	return currencyId, nil
}
