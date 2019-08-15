package main

import (
	"context"
	"log"

	"github.com/bookun/glue-sample/di"
)

func main() {
	app := di.NewApp()
	ctx := context.Background()
	if err := app.Start(ctx); err != nil {
		log.Fatal(ctx)
	}
}
