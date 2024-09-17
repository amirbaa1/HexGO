package repository

func (s *Repository) ConsumerEmail() ([]string, error) {
	queueName := "emailQueue"

	messages, err := s.rabbitMQClient.ConsumeMessages(queueName)
	var processedMessages []string

	if err != nil {
		return nil, nil
	}
	//forever := make(chan bool)

	for message := range messages {
		body := message.Body
		processedMessages = append(processedMessages, string(body))

	}

	//<-forever

	return processedMessages, nil
}
