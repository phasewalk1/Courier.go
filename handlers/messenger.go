package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/phasewalk1/courier-go/database"
	"github.com/phasewalk1/courier-go/extractors"
	"github.com/phasewalk1/courier-go/models"
	"gorm.io/gorm"
)

func HealthCheck(ctx *fiber.Ctx) error {
	return ctx.SendString("Hello, World!")
}

func CreateMessage(ctx *fiber.Ctx) error {
	message := new(models.Message)
	if err := ctx.BodyParser(message); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	database.DB.Db.Create(&message)

	return ctx.Status(200).JSON(message)
}

func GetMessageById(ctx *fiber.Ctx) error {
	messageID := ctx.Query("id")
	if messageID == "" {
		log.Println("Failed to parse message ID: ", messageID)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse message ID",
		})
	}

	message := new(models.Message)
	if err := database.DB.Db.First(&message, messageID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Message not found",
			})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(message)
}

func GetMessagesByUser(ctx *fiber.Ctx) error {
	params, perror := extractors.ExtractByUserParams(ctx)

	if perror != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": perror.Error(),
		})
	}

	if params.OnlySender() {
		messages, err := GetMessagesBySender(params.Sender)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return ctx.Status(fiber.StatusOK).JSON(messages)
	}

	if params.OnlyRecip() {
		messages, err := GetMessagesByRecip(params.Recip)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return ctx.Status(fiber.StatusOK).JSON(messages)
	}

	// TODO:
	// if params.Both() {}

	return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"error": "Cannot parse both sender and recip. Use /conv.",
	})
}

func GetMessagesBySender(sender string) ([]models.Message, error) {
	messages := []models.Message{}

	if err := database.DB.Db.Where("sender = ?", sender).Find(&messages).Error; err != nil {
		return nil, err
	}

	return messages, nil
}

func GetMessagesByRecip(recip string) ([]models.Message, error) {
	messages := []models.Message{}

	if err := database.DB.Db.Where("recipient = ?", recip).Find(&messages).Error; err != nil {
		return nil, err
	}

	return messages, nil
}
