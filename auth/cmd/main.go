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
	conn, err := server1.ConnectRabbit()

	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	rabbitMQClient, err := server1.NewRabbitMQ(conn)
	if err != nil {
		log.Fatalf("Failed to create RabbitMQ client: %v", err)
	}
	defer rabbitMQClient.Close()

	log.Println("Connected to RabbitMQ")

	db := server1.GetDB()
	userRepo := repository.NewRepository(db)

	userService := service.NewUserService(userRepo, rabbitMQClient)
	userHandler := handler.NewUserHandler(userService)

	userServer := server1.NewServer(userHandler)
	userServer.Initialize()
}
