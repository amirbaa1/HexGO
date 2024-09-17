package repository

import (
	"notfi/cmd/server"
	"notfi/internal/core/ports"
)

type Repository struct {
	rabbitMQClient server.RabbitMQClient
}

var _ ports.NotfiRepository = (*Repository)(nil)

func NewRepositoryNotfi(rabbitMQClient server.RabbitMQClient) *Repository {
	return &Repository{
		rabbitMQClient: rabbitMQClient,
	}
}
