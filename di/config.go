package di

import (
	"os"

	"go.uber.org/config"
	"go.uber.org/fx"
)

type Config struct {
	Env *config.YAML
}

func NewConfig() (*Config, error) {
	file, err := os.Open(".env")
	if err != nil {
		return nil, err
	}
	provider, err := config.NewYAML(config.Source(file))
	return &Config{Env: provider}, nil
}

var configfx = fx.Provide(NewConfig)
