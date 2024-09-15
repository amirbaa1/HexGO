package ports

import (
	"book/internal/core/model"
	"github.com/gofiber/fiber/v2"
)

type BookHandler interface {
	CreateBook(c *fiber.Ctx) error
	GetAllBooks(c *fiber.Ctx) error
	UpdateBook(c *fiber.Ctx) error
	GetBookById(c *fiber.Ctx) error
}

type BookService interface {
	CreateServiceBook(book *model.RequestBook) error
	GetAllBooks() (*[]model.ResponseBook, error)
	UpdateBook(book *model.RequestBookUpdate, bookId string) (*model.RequestBookUpdate, error)
	GetBookById(bookId string) (*model.ResponseBook, error)
}

type BookRepository interface {
	CreateBook(book *model.Book) error
	GetAllBooks() ([]model.Book, error)
	FindAuthorByFullNameForCreate(authorFirstName string, authorLastName string) (*model.Author, error)
	CreateAuthor(author *model.Author) error
	UpdateBook(book *model.Book) error
	GetBookById(id string) (*model.Book, error)
}
