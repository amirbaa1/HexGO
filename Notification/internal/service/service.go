package service

import "Notification/internal/core/ports"

type Service struct {
	message ports.MessagingPort
}

var _ ports.NotfiService = (*Service)(nil)

func NewService(message ports.MessagingPort) *Service {
	return &Service{
		message: message,
	}
}
