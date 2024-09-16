package service

import (
	"log"
)

func (s *Service) SendEmail() (bool, error) {
	queueName := "emailQueue"

	messages, err := s.rabbitMQClient.ConsumeMessages(queueName)
	if err != nil {
		return false, nil
	}
	for message := range messages {
		log.Println("[notfi] notfi service is running:", message.Body)
	}
	return true, nil
}
