package books

import (
	"github.com/gofiber/fiber/v2"
	"github.com/erica7dev/fiberapi/pkg/common/models"
)

type CreateBookRequestBody struct {
	Title string `json:"title"`
	Author string `json:"author"`
	Description string `json:"description"`
}

func(h handler) CreateBook(c *fiber.Ctx) error {
	body := CreateBookRequestBody{}

	//get request body
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var book models.Book

	book.Title = body.Title
	book.Author = body.Author
	book.Description = body.Description

	//save book
	h.DB.Create(&book)

	return c.Status(fiber.StatusCreated).JSON(&book)
}