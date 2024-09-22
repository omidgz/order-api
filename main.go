package main

import (
	"context"
	"log"

	"github.com/omidgz/order-api/app"
)

func main() {
	app := &app.App{}
	ctx := context.Background()
	if err := app.Start(ctx); err != nil {
		log.Fatal(err)
	}
}
