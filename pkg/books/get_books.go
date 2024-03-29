package books

import (
	"github.com/gofiber/fiber/v2"
	"github.com/erica7dev/fiberapi/pkg/common/models"
)

func (h handler) GetBooks(c *fiber.Ctx) error {
	var books []models.Book

	if result := h.DB.Find(&books); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&books)
}