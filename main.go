package main

import (
	"courses-api/src/config/builder"
	"log"
)

func main() {
	app, err := builder.BuildApp()
	if err != nil {
		log.Fatalf("Error iniciando la aplicación: %v", err)
	}

	router := app.GetRouter()
	router.Run(app.GetPort())
}
