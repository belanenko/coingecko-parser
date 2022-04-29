package main

import (
	"github.com/belanenko/coingecko-parser/internal/app/apiserver"
	"github.com/caarlos0/env"
	"github.com/sirupsen/logrus"
)

func main() {
	config := apiserver.NewConfig()
	if err := env.Parse(config); err != nil {
		logrus.Fatal(err)
	}

	apiserver := apiserver.New(config)
	if err := apiserver.Run(); err != nil {
		logrus.Fatal(err)
	}
}
