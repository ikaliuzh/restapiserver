package apiserver

import "http-rest-api/internal/app/storage"

type Config struct {
	BindAddr string `toml:"bind_addr"`
	LogLevel string `toml:"log_level"`
	Storage  *storage.Config
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
		Storage:  storage.NewConfig(),
	}
}
