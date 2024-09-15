package main

import (
	server2 "book/cmd/server"
	"book/internal/handler"
	"book/internal/repository"
	"book/internal/service"
)

func main() {
	server2.Connect()

	db := server2.GetDB()

	bookRepo := repository.NewRepository(db)
	bookService := service.NewServiceBook(bookRepo)
	bookHandler := handler.NewHandler(bookService)

	serverBook := server2.NewServer(bookHandler)
	serverBook.Internal()

}
