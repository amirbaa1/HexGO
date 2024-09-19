package ports

type NotfiService interface {
	SendEmail(message string) (bool, error)
}
type MessagingPort interface {
	ConsumeMessages(queueName string) (<-chan string, error)
	Close() error
}
