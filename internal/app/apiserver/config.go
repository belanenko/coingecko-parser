package apiserver

import (
	"github.com/belanenko/coingecko-parser/internal/app/storage"
)

type Config struct {
	BindAddress    string `env:"BIND_ADDRESS"`
	LogLevel       string `env:"LOG_LEVEL"`
	WalletsHistory *storage.WalletsHistory
}

func NewConfig() *Config {
	return &Config{
		BindAddress:    ":8080",
		LogLevel:       "debug",
		WalletsHistory: storage.NewWalletsHistory(),
	}
}
