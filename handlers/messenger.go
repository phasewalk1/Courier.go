package handlers

import (
	"github.com/phasewalk1/courier-go/models"
	"github.com/phasewalk1/courier-go/database"
	"github.com/gofiber/fiber/v2"
)

func HealthCheck(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func CreateMessage(c *fiber.Ctx) error {
	message := new(models.Message)
	if err := c.BodyParser(message); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map {
			"error": err.Error(),
		})
	}

	database.DB.Db.Create(&message)

	return c.Status(200).JSON(message)
}
