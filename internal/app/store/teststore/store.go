package teststore

import "github.com/belanenko/coingecko-parser/internal/app/model"

type Store struct {
	db                map[string][]model.History
	historyRepository *HistoryRepository
}

func New() *Store {
	return &Store{
		db: make(map[string][]model.History),
	}
}

// Игнорирую использование контекста/потом допилю

func (s *Store) History() *HistoryRepository {
	if s.historyRepository != nil {
		return s.historyRepository
	}
	s.historyRepository = &HistoryRepository{
		Store: s,
	}
	return s.historyRepository
}
