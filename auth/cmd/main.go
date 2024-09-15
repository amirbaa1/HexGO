package main

import (
	server2 "auth/cmd/server"
	"auth/internal/handler"
	"auth/internal/repository"
	"auth/internal/service"
)

func main() {
	server2.Connect()

	db := server2.GetDB()
	userRepo := repository.NewRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	userServer := server2.NewServer(userHandler)
	userServer.Initialize()
}
