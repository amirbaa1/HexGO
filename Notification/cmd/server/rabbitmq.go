package server

import (
	"Notification/internal/core/ports"
	ampq "github.com/rabbitmq/amqp091-go"
	"log"
)

type RabbitMQClient struct {
	conn *ampq.Connection
	ch   *ampq.Channel
}

var _ ports.MessagingPort = (*RabbitMQClient)(nil)

func ConnectRabbit() (*ampq.Connection, error) {
	return ampq.Dial("amqp://guest:guest@localhost:5672/")
}

func NewRabbitMQ(conn *ampq.Connection) (RabbitMQClient, error) {
	ch, err := conn.Channel()
	if err != nil {
		return RabbitMQClient{}, err
	}

	return RabbitMQClient{
		conn: conn,
		ch:   ch,
	}, nil
}

func (rc RabbitMQClient) Close() error {
	return rc.ch.Close()
}

func (rc RabbitMQClient) CreateQueueDeclare(queueName string, durable, autodelete bool) error {
	_, err := rc.ch.QueueDeclare(queueName,
		durable,
		autodelete,
		false,
		false,
		nil)

	return err
}

func (rc RabbitMQClient) ConsumeMessages(queueName string) (<-chan string, error) {
	messages, err := rc.ch.Consume(
		queueName,
		"",
		true,  // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,
	)
	if err != nil {
		log.Printf("Failed to consume messages from queue %s: %s", queueName, err)
		return nil, err
	}
	log.Printf("Successfully consumed messages from queue %s", queueName)

	msgChan := make(chan string)

	go func() {
		for message := range messages {
			msgChan <- string(message.Body)
		}
	}()

	return msgChan, nil
}
