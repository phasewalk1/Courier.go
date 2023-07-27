package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/phasewalk1/courier-go/handlers"
)

// @root
// /api/v1
func BuildRouter(v1 fiber.Router) {
	// @health
	v1.Get("/", handlers.HealthCheck)
	// @data
	// TransientMessage
	v1.Post("/messages/", handlers.CreateMessage)
	// @queryparams
	// id: int
	v1.Get("/messages/one/", handlers.GetMessageById)
	// @queryparams
	// ?sender: string
	// ?recip: string
	v1.Get("/messages/", handlers.GetMessagesByUser)
}
