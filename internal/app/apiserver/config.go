package apiserver

import "github.com/belanenko/coingecko-parser/internal/app/store"

type Config struct {
	Store       *store.Config
	BindAddress string `env:"BIND_ADDRESS"`
}

func NewConfig(store *store.Config) *Config {
	return &Config{Store: store}
}
