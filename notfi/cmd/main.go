package main

import (
	"log"
	"notfi/cmd/server"
	"notfi/internal/service"
)

func main() {
	//conn, err := server.ConnectRabbit()
	//if err != nil {
	//	log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	//}
	//defer conn.Close()
	//
	//rabbitMQClient, err := server.NewRabbitMQ(conn)
	//if err != nil {
	//	log.Fatalf("Failed to create RabbitMQ client: %v", err)
	//}
	//defer rabbitMQClient.Close()
	//
	////repository1 := repository.NewRepositoryNotfi(rabbitMQClient)
	//ss := service.NewService(rabbitMQClient)
	//log.Println("Starting server", ss)
	//
	//log.Println("Connected to RabbitMQ")

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

	messages, err := rabbitMQClient.ConsumeMessages("emailQueue")
	if err != nil {
		log.Fatalf("Failed to create RabbitMQ client: %v", err)
	}

	log.Println("Connected to RabbitMQ")
	forever := make(chan bool)

	emailService := service.NewService(rabbitMQClient)

	go func() {
		for message := range messages {
			log.Printf("Received a message: %s", message.Body)

			if send, err := emailService.SendEmail(string(message.Body)); err != nil {
				log.Printf("Failed to send email: %v", err)
			} else {
				log.Println("Email sent successfully", send)
			}
		}
	}()

	log.Println("Waiting for messages. To exit press CTRL+C")
	<-forever

	server.Internal()

}
