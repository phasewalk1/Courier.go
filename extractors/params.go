package extractors

import (
	"github.com/gofiber/fiber/v2"
	"github.com/phasewalk1/courier-go/models"
)

func ExtractByUserParams(ctx *fiber.Ctx) (models.ByUserSearch, error) {
	paramsNull := models.ByUserSearch{}
	s, r := ctx.Query("sender"), ctx.Query("recip")
	if s == "" && r == "" {
		return paramsNull, ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse any params",
		})
	}

	params := models.ByUserSearch{Sender: s, Recip: r}
	return params, nil
}
