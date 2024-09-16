package server

import (
	"auth/internal/core/ports"
	"fmt"
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
	if queueName == "" {
		return fmt.Errorf("queue name is empty")
	}

	_, err := rc.ch.QueueDeclare(
		queueName,  // name of the queue
		durable,    // durable
		autodelete, // delete when unused
		false,      // exclusive
		false,      // no-wait
		nil,        // arguments
	)

	if err != nil {
		log.Printf("Failed to declare queue: %s", err)
	}
	return err
}

func (rc RabbitMQClient) PublishMessage(queueName, message string) error {
	if queueName == "" {
		return fmt.Errorf("queue name is empty")
	}

	err := rc.ch.Publish(
		"",        // exchange
		queueName, // routing key (queue name)
		false,     // mandatory
		false,     // immediate
		ampq.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)

	if err != nil {
		log.Printf("Failed to publish message to queue %s: %s", queueName, err)
	}
	return err

}
