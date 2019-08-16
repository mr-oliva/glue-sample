package di

import (
	"fmt"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Logger struct {
	Logger *zap.Logger
}

func NewLogger(c *Config) (*Logger, error) {
	if c.Yaml.Get("APP_ENV").String() == "production" {
		l, err := zap.NewProduction()
		if err != nil {
			return nil, err
		}
		return &Logger{l}, nil
	}

	l, err := zap.NewDevelopment()
	if err != nil {
		return nil, err
	}
	return &Logger{l}, nil
}

func (l *Logger) Printf(format string, msg ...interface{}) {
	l.Logger.Info(fmt.Sprintf(format, msg...))
}

func (l *Logger) Errorf(format string, msg ...interface{}) {
	l.Logger.Error(fmt.Sprintf(format, msg...))
}

var loggerfx = fx.Provide(NewLogger)
