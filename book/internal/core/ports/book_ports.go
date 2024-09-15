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
	GetAllBooks() (*[]model.ResponseBook, error)
}

type BookRepository interface {
	CreateBook(book *model.Book) error
	GetAllBooks() ([]model.Book, error)
	FindAuthorByFullNameForCreate(authorFirstName string, authorLastName string) (*model.Author, error)
	CreateAuthor(author *model.Author) error
}
