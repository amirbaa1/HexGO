package server

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"Notification/internal/service"
)

func Internal() {
	app := fiber.New()
	app.Use(logger.New())

	err := app.Listen(":3033")
	if err != nil {
		log.Fatal(err)
	}
}

func StartRabbitMQ() {
	//start rabbitmq
	conn, err := ConnectRabbit()
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	rabbitMQClient, err := NewRabbitMQ(conn)
	if err != nil {
		log.Fatalf("Failed to create RabbitMQ client: %v", err)
	}
	defer rabbitMQClient.Close()
	log.Println("1 Connected to RabbitMQ")
	messages, err := rabbitMQClient.ConsumeMessages("emailQueue")
	if err != nil {
		log.Fatalf("Failed to consume messages from RabbitMQ: %v", err)
	}

	log.Println("2 Connected to RabbitMQ")

	emailService := service.NewService(rabbitMQClient)

	go func() {
		for message := range messages {
			log.Printf("Received a message: %s", message)

			if send, err := emailService.SendEmail(message); err != nil {
				log.Printf("Failed to send email: %v", err)
			} else {
				log.Println("Email sent successfully", send)
			}
		}
	}()

	log.Println("Waiting for messages. To exit press CTRL+C")

	select {}
}
