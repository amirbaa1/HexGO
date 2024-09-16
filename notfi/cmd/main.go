package main

import (
	"log"
	"notfi/cmd/server"
	"notfi/internal/service"
)

func main() {
	conn, err := server.ConnectRabbit()
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	rabbitMQClient, err := server.NewRabbitMQ(conn)
	if err != nil {
		log.Fatalf("Failed to create RabbitMQ client: %v", err)
	}
	defer rabbitMQClient.Close()

	service.NewNotfiService(rabbitMQClient)

	log.Println("Connected to RabbitMQ")

	server.Internal()

}
