package service

import (
	"notfi/cmd/server"
	"notfi/internal/core/ports"
)

type Service struct {
	rabbitMQClient server.RabbitMQClient
}

var _ ports.NotfiService = (*Service)(nil)

func NewNotfiService(rabbitMQClient server.RabbitMQClient) *Service {
	return &Service{
		rabbitMQClient: rabbitMQClient,
	}
}
