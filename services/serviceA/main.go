package main

import (
	"context"
	"log"
	"net/http"

	"github.com/bookun/glue-sample/di"
	"github.com/bookun/glue-sample/services/serviceA/controllers"
	"github.com/bookun/glue-sample/services/serviceA/gateways"
	"github.com/bookun/glue-sample/services/serviceA/handlers"
	"github.com/bookun/glue-sample/services/serviceA/repositories"
	"go.uber.org/fx"
)

func NewRouter(userHandler *handlers.User, ipHandler *handlers.IP) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/users", userHandler.HandleTaskGetUsers)
	mux.HandleFunc("/user/", userHandler.HandleTaskGetUserNameById)
	mux.HandleFunc("/ip", ipHandler.HandleTaskGetMyGIP)
	return mux
}

func StartServer(mux *http.ServeMux, c *di.Config) error {
	if err := http.ListenAndServe(":"+c.Get("PORT"), mux); err != nil {
		return err
	}
	return nil
}

func main() {
	app := fx.New(
		di.Module,
		fx.Provide(
			func() string {
				return "testdata/user.csv"
			},
			repositories.NewFile,
			gateways.NewGIPGateway,
			controllers.NewUser,
			controllers.NewIPGateway,
			handlers.NewUser,
			handlers.NewIP,
			NewRouter,
		),
		fx.Invoke(StartServer),
	)
	ctx := context.Background()
	if err := app.Start(ctx); err != nil {
		log.Fatal(ctx)
	}
}
