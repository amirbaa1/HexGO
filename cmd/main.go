package main

import (
	"HexGO/internal/handler"
	"HexGO/internal/repository"
	"HexGO/internal/service"
	"HexGO/server"
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
