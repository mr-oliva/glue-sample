package di

import (
	"context"
	"net/http"
	"os"

	"go.uber.org/fx"
)

func NewMux(lc fx.Lifecycle, logger *Logger) *http.ServeMux {
	logger.Printf("Executing NewMux.")
	mux := http.NewServeMux()
	server := &http.Server{
		Addr:    ":" + os.Getenv("PORT"),
		Handler: mux,
	}
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			logger.Printf("Starting HTTP server.")
			go server.ListenAndServe()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Printf("Stopping HTTP server.")
			return server.Shutdown(ctx)
		},
	})

	return mux
}

var serverfx = fx.Provide(NewMux)
