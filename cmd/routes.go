package main

import (
	"github.com/phasewalk1/courier-go/handlers"
	"github.com/gofiber/fiber/v2"
)

func initRouter(app *fiber.App) {
	app.Get("/", handlers.HealthCheck)
	app.Post("/api/v1/messages/", handlers.CreateMessage)
}
