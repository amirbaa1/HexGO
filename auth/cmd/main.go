package main

import (
	"auth/internal/handler"
	"auth/internal/repository"
	"auth/internal/service"
	"auth/server"
)

func main() {
	server.Connect()

	db := server.GetDB()
	userRepo := repository.NewRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	userServer := server.NewServer(userHandler)
	userServer.Initialize()
}
