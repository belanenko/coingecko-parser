package main

import (
	"log"
	"os"
	"strings"

	"github.com/belanenko/coingecko-parser/internal/app/apiserver"
	"github.com/belanenko/coingecko-parser/internal/app/geckoparser"
	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	apiserverConfig := apiserver.NewConfig()
	env.Parse(apiserverConfig)

	currencies := strings.Split(os.Getenv("WALLETS"), ",")

	for i := range currencies {
		currencies[i] = strings.TrimSpace(currencies[i])
	}

	gecko := geckoparser.New(currencies)

	if err := apiserver.Start(apiserverConfig, gecko); err != nil {
		log.Fatal(err)
	}
}
