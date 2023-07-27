package main

import (
	"log"
	"github.com/gofiber/fiber/v2"

	"github.com/phasewalk1/courier-go/database"
)

// main function
func main() {
	// connect postgres
	database.ConnectDb()
    // create app
    app := fiber.New()
	
	// setup the router
	initRouter(app)

    log.Fatal(app.Listen(":3000"))
}
