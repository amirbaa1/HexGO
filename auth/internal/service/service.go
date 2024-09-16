package service

import (
	"auth/internal/core/ports"
)

type UserService struct {
	userRepository ports.UserRepository
	message        ports.MessagingPort
}

var _ ports.UserService = (*UserService)(nil)

func NewUserService(userRepository ports.UserRepository, message ports.MessagingPort) *UserService {
	return &UserService{
		userRepository: userRepository,
		message:        message,
	}
}
