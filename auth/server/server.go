package server

import (
	"auth/internal/core/ports"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type UserServer struct {
	userHandler ports.UserHandler
}

func NewServer(userHandler ports.UserHandler) *UserServer {
	return &UserServer{
		userHandler: userHandler,
	}
}

func (s *UserServer) Initialize() {
	app := fiber.New()
	app.Use(logger.New())

	v1 := app.Group("/v1")

	userRoute := v1.Group("/user")
	userRoute.Post("/Register", s.userHandler.Register)
	userRoute.Post("/Login", s.userHandler.Login)

	err := app.Listen(":3000")
	if err != nil {
		log.Fatal(err)
	}
}
