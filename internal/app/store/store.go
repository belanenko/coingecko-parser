package store

import (
	"context"

	"github.com/jackc/pgx/v4"
)

type Store struct {
	config            *Config
	db                *pgx.Conn
	historyRepository *HistoryRepository
}

func NewStore(config *Config) *Store {
	return &Store{
		config: config,
	}
}

func (s *Store) Open() error {
	db, err := pgx.Connect(context.Background(), s.config.DatabaseURL)
	if err != nil {
		return err
	}
	if err := db.Ping(context.Background()); err != nil {
		return err
	}

	s.db = db
	return nil
}

func (s *Store) Close() {
	s.db.Close(context.Background())
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
