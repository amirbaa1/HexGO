package ports

import (
	"book/internal/core/model"
	"github.com/gofiber/fiber/v2"
)

type BookHandler interface {
	CreateBook(c *fiber.Ctx) error
	GetAllBooks(c *fiber.Ctx) error
}

type BookService interface {
	CreateServiceBook(book *model.RequestBook) error
	GetAllBooks(book *model.RequestBook) error
}

type BookRepository interface {
	CreateBook(book *model.Book) (string, error)
	GetAllBooks(book model.Book) ([]string, error) // model.Response ?
}
