package main

import (
	"context"
	"github.com/pkg/errors"
	"log"
	"measurements-api/internal/app"
)

func main() {
	application, err := app.NewApp(context.Background())
	if err != nil {
		log.Fatal(errors.Wrap(err, "application initialization failed"))
	}

	err = application.Run()
	if err != nil {
		log.Fatal(errors.Wrap(err, "application run failed"))
	}
}
