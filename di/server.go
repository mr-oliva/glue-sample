package di

import (
	"net/http"
	"os"

	"go.uber.org/zap"
)

type Server struct {
	Logger *zap.Logger
}

func NewServer(logger *zap.Logger) *Server {
	return &Server{Logger: logger}
}

func StartServer(mux *http.ServeMux) error {
	if err := http.ListenAndServe(":"+os.Getenv("PORT"), mux); err != nil {
		return err
	}
	return nil
}
