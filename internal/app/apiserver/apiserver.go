package apiserver

import (
	"context"
	"net/http"

	"github.com/belanenko/coingecko-parser/internal/app/geckoparser"
	"github.com/belanenko/coingecko-parser/internal/app/store/sqlstore"
	"github.com/jackc/pgx/v4"
)

func Start(config *Config, gecko *geckoparser.GeckoParser) error {
	db, err := newDB(config.DatabaseURL)
	if err != nil {
		return err
	}
	defer db.Close(context.Background())
	store := sqlstore.New(db)
	srv := newServer(store, gecko)
	if err := srv.parseCurrenciesHistory(); err != nil {
		return err
	}

	return http.ListenAndServe(config.BindAddress, srv)
}

func newDB(databaseURL string) (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), databaseURL)
	if err != nil {
		return nil, err
	}
	if err := conn.Ping(context.Background()); err != nil {
		return nil, err
	}
	return conn, nil
}
