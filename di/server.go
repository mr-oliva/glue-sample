package di

import "go.uber.org/fx"

type Server struct {
	Logger *Logger
}

func NewServer(logger *Logger) *Server {
	return &Server{Logger: logger}
}

var serverfx = fx.Provide(NewServer)
