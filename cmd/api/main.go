package main

import (
	"log"

	"fds-backend/internal/app"
	"fds-backend/internal/bootstrap"
)

func main() {
	application, err := bootstrap.BuildApplication()
	if err != nil {
		log.Fatal(err)
	}

	if err := app.Run(application); err != nil {
		log.Fatal(err)
	}
}
