package storage

import (
	"sync"

	"github.com/belanenko/coingecko-parser/internal/app/model"
)

type WalletsHistory struct {
	sync.RWMutex
	list map[string][]model.PricePerTimestamp
}

func NewWalletsHistory() *WalletsHistory {
	walletsMap := make(map[string][]model.PricePerTimestamp, 20)

	return &WalletsHistory{
		list: walletsMap,
	}
}

func (h *WalletsHistory) AddHistory(wallet string, ppt []model.PricePerTimestamp) {
	h.Lock()
	defer h.Unlock()

	for _, p := range ppt {
		h.list[wallet] = append(h.list[wallet], p)
	}
}

func (h *WalletsHistory) GetHistoryForWallet(wallet string) []model.PricePerTimestamp {
	h.RLock()
	defer h.RUnlock()
	return h.list[wallet]
}

func (h *WalletsHistory) GetAllHistory() *map[string][]model.PricePerTimestamp {
	h.RLock()
	defer h.RUnlock()
	return &h.list
}
