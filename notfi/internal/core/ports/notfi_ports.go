package ports

import ampq "github.com/rabbitmq/amqp091-go"

type NotfiService interface {
	SendEmail(message string) (bool, error)
}

type NotfiRepository interface {
	ConsumerEmail() ([]string, error)
}
type MessagingPort interface {
	ConsumeMessages(queueName string) (<-chan ampq.Delivery, error)
}
