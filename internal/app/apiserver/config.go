package apiserver

type Config struct {
	BindAddress string `env:"BIND_ADDRESS"`
	DatabaseURL string `env:"DATABASE_URL"`
}

func NewConfig() *Config {
	return &Config{}
}
