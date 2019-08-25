package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/bookun/glue-sample/di"
	"github.com/bookun/glue-sample/services/serviceA/controllers"
	"github.com/bookun/glue-sample/services/serviceA/gateways"
	"github.com/bookun/glue-sample/services/serviceA/handlers"
	"github.com/bookun/glue-sample/services/serviceA/repositories"
	"go.uber.org/fx"
)

func Register(mux *http.ServeMux, userHandler *handlers.User, ipHandler *handlers.IP) {
	mux.HandleFunc("/users", userHandler.HandleTaskGetUsers)
	mux.HandleFunc("/user/", userHandler.HandleTaskGetUserNameById)
	mux.HandleFunc("/ip", ipHandler.HandleTaskGetMyGIP)
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
		),
		fx.Invoke(Register),
	)
	ctx := context.Background()
	if err := app.Start(ctx); err != nil {
		log.Fatal(err)
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	<-quit

	if err := app.Stop(ctx); err != nil {
		log.Fatal(err)
	}
}
