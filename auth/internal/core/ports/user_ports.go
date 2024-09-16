package ports

import (
	"auth/internal/core/model"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type UserHandler interface {
	Register(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
	Profile(c *fiber.Ctx) error
}

type UserRepository interface {
	Register(auth *model.User) error
	//Login(login *model.AuthRequest) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
}
type UserService interface {
	Register(register *model.RegisterRequest) error
	Login(login *model.AuthRequest) (model.AuthResponse, error)
	Profile(token *jwt.Token) (model.ProfileResponse, error)
	SendEmail(sm *model.EmailMessage) (bool, error)
}

type MessagingPort interface {
	PublishMessage(queueName string, message string) error
}

type UserServer interface {
	Initialize()
}
