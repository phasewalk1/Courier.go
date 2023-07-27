package middleware

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func ApplyCORS(app *fiber.App) {
	sv := os.Getenv("AUTHORIZED_SERVICE")
	app.Use(cors.New(cors.Config {
		AllowOrigins: sv,
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
}
