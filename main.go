package main

import (
	"courses-api/src/config/builder"
)

func main() {
	app := builder.BuildApp()
	router := app.GetRouter()
	router.Run(app.GetPort())
}
