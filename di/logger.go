package di

import "go.uber.org/zap"

type Logger struct {
	Logger *zap.Logger
}

func NewLogger(c Config) (*Logger, error) {
	if c.Env != "production" {
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

func (l *Logger) Print(label string, msg interface{}) {
	l.Logger.Info(label, zap.Any(label, msg))
}

func (l *Logger) Error(label string, msg interface{}) {
	l.Logger.Error(label, zap.Any(label, msg))
}
