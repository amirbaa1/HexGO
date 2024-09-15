package main

import (
	"book/internal/handler"
	"book/internal/repository"
	"book/internal/service"
	"book/server"
)

func main() {
	server.Connect()

	db := server.GetDB()

	bookRepo := repository.NewRepository(db)
	bookService := service.NewServiceBook(bookRepo)
	bookHandler := handler.NewHandler(bookService)

	serverBook := server.NewServer(bookHandler)
	serverBook.Internal()

}
