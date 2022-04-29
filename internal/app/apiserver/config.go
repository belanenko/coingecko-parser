package apiserver

type Config struct {
	BindAddress string `env:"BIND_ADDRESS"`
	LogLevel    string `env:"LOG_LEVEL"`
}

func NewConfig() *Config {
	return &Config{
		BindAddress: ":8080",
		LogLevel:    "debug",
	}
}
