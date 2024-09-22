package main

import (
	"context"
	"log"

	"github.com/omidgz/order-api/app"
)

func main() {
	app := app.New()
	ctx := context.Background()
	if err := app.Start(ctx); err != nil {
		log.Fatal(err)
	}
}
