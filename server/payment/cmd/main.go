package main

import (
	"context"
	"log"

	"github.com/omaqase/sato/payment/internal/di"
)

func main() {
	app := di.NewContainer()

	if err := app.Start(context.Background()); err != nil {
		log.Fatal(err)
	}

	<-app.Done()
}
