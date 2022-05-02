package main

import (
	"log"
	"os"
	"strings"

	"github.com/belanenko/coingecko-parser/internal/app/apiserver"
	"github.com/belanenko/coingecko-parser/internal/app/geckoparser"
	"github.com/belanenko/coingecko-parser/internal/app/store"
	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	storeConfig := store.NewConfig()
	env.Parse(storeConfig)

	apiserverConfig := apiserver.NewConfig(storeConfig)
	env.Parse(apiserverConfig)

	wallets := strings.Split(os.Getenv("WALLETS"), ",")

	for i := range wallets {
		wallets[i] = strings.TrimSpace(wallets[i])
	}

	gecko := geckoparser.New(wallets)

	apiserver := apiserver.New(apiserverConfig, gecko)
	if err := apiserver.Start(); err != nil {
		log.Fatal(err)
	}
}
