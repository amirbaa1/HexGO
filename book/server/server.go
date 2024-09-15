package server

import (
	"book/internal/core/ports"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Server struct {
	bookHandler ports.BookHandler
}

func NewServer(bookHandler ports.BookHandler) *Server {
	return &Server{
		bookHandler: bookHandler,
	}
}

func (ser *Server) Internal() {
	app := fiber.New()
	app.Use(logger.New())

	v1 := app.Group("/v1")
	bookRoute := v1.Group("/book")
	bookRoute.Get("/", ser.bookHandler.GetAllBooks)
	bookRoute.Post("/AddBook", ser.bookHandler.CreateBook)
	bookRoute.Put("/UpdateBook/:id", ser.bookHandler.UpdateBook)
	bookRoute.Get("/book/:id", ser.bookHandler.GetBookById)

	err := app.Listen(":3001")
	if err != nil {
		log.Fatal(err)
	}

}
