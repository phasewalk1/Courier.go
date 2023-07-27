package main

import (
	"github.com/gofiber/fiber/v2"
	"log"

	"github.com/phasewalk1/courier-go/database"
	"github.com/phasewalk1/courier-go/router"
)

// main function
func main() {
	// connect postgres
	database.ConnectDb()
	// create app
	app := fiber.New()

	api := app.Group("/api")
	v1 := api.Group("/v1")

	// setup the routes
	router.BuildRouter(v1)

	log.Fatal(app.Listen(":3000"))
}
