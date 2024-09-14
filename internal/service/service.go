package service

import (
	"HexGO/internal/core/ports"
)

type UserService struct {
	userRepository ports.UserRepository
}

var _ ports.UserService = (*UserService)(nil)

func NewUserService(userRepository ports.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}
