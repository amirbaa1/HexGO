package main

import (
	server1 "auth/cmd/server"
	"auth/internal/handler"
	"auth/internal/repository"
	"auth/internal/service"
	"go.uber.org/zap"
)

func main() {

	server1.Connect()

	logger, _ := zap.NewProduction()

	defer logger.Sync()

	conn, err := server1.ConnectRabbit()

	if err != nil {
		logger.Fatal("Failed to connect to RabbitMQ", zap.Error(err))
	}
	defer conn.Close()

	rabbitMQClient, err := server1.NewRabbitMQ(conn)
	if err != nil {
		logger.Fatal("Failed to create RabbitMQ client: %v", zap.Error(err))
	}
	defer rabbitMQClient.Close()

	logger.Info("Connected to RabbitMQ")

	err = rabbitMQClient.CreateQueueDeclare("emailQueue", true, false)
	if err != nil {
		logger.Fatal("Failed to create queue: %v", zap.Error(err))
	}

	db := server1.GetDB()
	userRepo := repository.NewRepository(db)

	userService := service.NewUserService(userRepo, rabbitMQClient)
	userHandler := handler.NewUserHandler(userService)

	userServer := server1.NewServer(userHandler)
	userServer.Initialize()
}
