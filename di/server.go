package di

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"

	"go.uber.org/fx"
)

func NewMux(lc fx.Lifecycle, logger *Logger) *http.ServeMux {
	logger.Printf("Executing NewMux.")
	mux := http.NewServeMux()
	host := os.Getenv("HOST")
	if host == "" {
		//TODO: later
		host, _ = getPrivateIp()
	}
	server := &http.Server{
		Addr:    host + ":" + os.Getenv("PORT"),
		Handler: mux,
	}
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			logger.Printf("Starting HTTP server.")
			logger.Printf("Server Addr %s\n", "http://"+host+":"+os.Getenv("PORT"))
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

func getPrivateIp() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(), nil
			}
		}
	}
	return "", fmt.Errorf("failed to find local ip addr")
}

var serverfx = fx.Provide(NewMux)
