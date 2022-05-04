package main

import (
	"log"
	"os"
	"strings"

	"github.com/belanenko/coingecko-parser/internal/app/apiserver"
	"github.com/belanenko/coingecko-parser/internal/app/parser/geckoparser"
	"github.com/belanenko/coingecko-parser/internal/app/store/sqlstore"
	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	apiserverConfig := apiserver.NewConfig()
	env.Parse(apiserverConfig)

	migrator := sqlstore.NewMigrator(apiserverConfig.DatabaseURL)
	if err := migrator.Up(); err != nil {
		log.Fatal(err)
	}

	currencies := strings.Split(os.Getenv("WALLETS"), ",")

	for i := range currencies {
		currencies[i] = strings.TrimSpace(currencies[i])
	}

	gecko := geckoparser.New(currencies)

	if err := apiserver.Start(apiserverConfig, gecko); err != nil {
		log.Fatal(err)
	}
}
