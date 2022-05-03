package sqlstore

import (
	"github.com/belanenko/coingecko-parser/internal/app/store"
	"github.com/jackc/pgx/v4"
)

type Store struct {
	db                *pgx.Conn
	historyRepository *HistoryRepository
}

func New(db *pgx.Conn) *Store {
	return &Store{
		db: db,
	}
}

// Игнорирую использование контекста/потом допилю

func (s *Store) History() store.HistoryRepository {
	if s.historyRepository != nil {
		return s.historyRepository
	}
	s.historyRepository = &HistoryRepository{
		Store: s,
	}
	return s.historyRepository
}
