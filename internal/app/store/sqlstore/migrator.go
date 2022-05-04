package sqlstore

import (
	"database/sql"
	"embed"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

const MigrationsTableName = "migrations"

var embedMigrations embed.FS

type Migrator struct {
	databaseURL string
}

func NewMigrator(databaseURL string) *Migrator {
	return &Migrator{
		databaseURL: databaseURL,
	}
}

func (m *Migrator) Up() error {
	db, err := sql.Open("postgres", m.databaseURL)
	if err != nil {
		return fmt.Errorf("open database connection error: %v", err)
	}
	defer db.Close()

	//goose.SetBaseFS(embedMigrations)
	goose.SetTableName(MigrationsTableName)
	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}

	if err := goose.Up(db, "migrations"); err != nil {
		return fmt.Errorf("migrate error: %w", err)
	}

	return nil
}

func (m *Migrator) Down() error {
	db, err := sql.Open("postgres", m.databaseURL)
	if err != nil {
		return fmt.Errorf("open database connection error: %v", err)
	}
	defer db.Close()

	//goose.SetBaseFS(embedMigrations)
	goose.SetTableName(MigrationsTableName)
	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}

	if err := goose.Down(db, "migrations"); err != nil {
		return fmt.Errorf("migrate error: %w", err)
	}

	return nil
}
