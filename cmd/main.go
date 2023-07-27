package main

import (
	"github.com/gofiber/fiber/v2"
	"log"

	"github.com/phasewalk1/courier-go/database"
	"github.com/phasewalk1/courier-go/router"
	"github.com/phasewalk1/courier-go/middleware"
)

func main() {
	database.ConnectDb()

	app := fiber.New()

	middleware.ApplyCORS(app)
	middleware.ApplyLimiter(app)
	middleware.AttachLogger(app)

	api := app.Group("/api")
	v1 := api.Group("/v1")

	router.BuildRouter(v1)

	log.Fatal(app.Listen(":3000"))
}
