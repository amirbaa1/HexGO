package ports

import (
	"auth/internal/core/model"
	"github.com/golang-jwt/jwt/v5"
)

type UserService interface {
	Register(register *model.RegisterRequest) error
	Login(login *model.AuthRequest) (model.AuthResponse, error)
	Profile(token *jwt.Token) (model.ProfileResponse, error)
}
