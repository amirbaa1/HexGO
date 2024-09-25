package main

import (
	server1 "auth/cmd/server"
	"auth/internal/handler"
	"auth/internal/repository"
	"auth/internal/service"
	"log"
)

func main() {

	server1.Connect()

	rabbitClient, err := server1.StartRabbitMQ()
	if err != nil {
		log.Fatalf("Failed to start RabbitMQ: %v", err)
	}
	
	db := server1.GetDB()
	userRepo := repository.NewRepository(db)

	userService := service.NewUserService(userRepo, rabbitClient)
	userHandler := handler.NewUserHandler(userService)

	userServer := server1.NewServer(userHandler)
	userServer.Initialize()
}
