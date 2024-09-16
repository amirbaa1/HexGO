package server

import (
	ampq "github.com/rabbitmq/amqp091-go"
	"log"
)

type RabbitMQClient struct {
	conn *ampq.Connection
	ch   *ampq.Channel
}

func ConnectRabbit() (*ampq.Connection, error) {
	return ampq.Dial("amqp://guest:guest@localhost:5672/")
}

func NewRabbitMQ(conn *ampq.Connection) (RabbitMQClient, error) {
	ch, err := conn.Channel()
	if err != nil {
		return RabbitMQClient{}, err
	}

	//ch.QueueDeclare()
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

// func (rc RabbitMQClient) PublishMessage(queueName, message string) error {
// 	err := rc.ch.Publish(
// 		"",
// 		queueName,
// 		false,
// 		false,
// 		ampq.Publishing{
// 			ContentType: "text/plain",
// 			Body:        []byte(message),
// 		})

// 	if err != nil {
// 		log.Printf("Failed to publish message to queue %s: %s", queueName, err)
// 	}
// 	return err

// }
func (rc RabbitMQClient) ConsumeMessages(queueName string) (<-chan ampq.Delivery, error) {
	message, err := rc.ch.Consume(
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
	return message, nil
}
