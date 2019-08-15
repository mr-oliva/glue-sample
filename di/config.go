package di

import (
	"os"

	"go.uber.org/fx"
)

type Config struct {
	Env string
}

func NewConfig() *Config {
	if os.Getenv("APP_ENV") == "" {
		return &Config{Env: "development"}
	}
	return &Config{Env: os.Getenv("APP_ENV")}
}

func apply(*fx.App) *Config {
	if os.Getenv("APP_ENV") == "" {
		return &Config{Env: "development"}
	}
	return &Config{Env: os.Getenv("APP_ENV")}
}
