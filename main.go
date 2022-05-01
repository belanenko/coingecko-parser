package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/belanenko/coingecko-parser/internal/app/apiserver"
	"github.com/belanenko/coingecko-parser/internal/app/gecko"
	"github.com/belanenko/coingecko-parser/internal/app/storage"
	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	godotenv.Load()

	wallets := strings.Split(os.Getenv("WALLETS"), ",")
	if len(wallets) == 0 || len(wallets) > 50 {
		log.Fatal("Кол-во валют должно быть больше 0 и меньше 50.")
	}

	history := storage.NewWalletsHistory()
	for _, wallet := range wallets {
		p, err := gecko.GetPriceForLasyYear(wallet)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: %v\n", wallet, err)
			continue
		}
		history.AddHistory(wallet, p)
	}

	config := apiserver.NewConfig()
	if err := env.Parse(config); err != nil {
		logrus.Fatal(err)
	}
	config.WalletsHistory = history

	apiserver := apiserver.New(config)
	if err := apiserver.Run(); err != nil {
		logrus.Fatal(err)
	}
}
