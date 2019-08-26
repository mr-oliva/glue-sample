package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"go.uber.org/fx"
)

type Process struct {
	name string
}

func NewProcess(lc fx.Lifecycle, name string) *Process {
	process := &Process{name}
	lc.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				log.Printf("%s is generated\n", process.name)
				go func() {
					for {
						log.Printf("%s is living\n", process.name)
						time.Sleep(1 * time.Second)
					}
				}()
				return nil
			},
			OnStop: func(ctx context.Context) error {
				fmt.Printf("%s is killed\n", process.name)
				return nil
			},
		},
	)
	return process
}

func Show(p *Process) {
	fmt.Printf("name: %s\n", p.name)
}

func main() {
	app := fx.New(
		fx.Provide(
			func() string { return "processA" },
			NewProcess,
		),
		fx.Invoke(Show),
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
