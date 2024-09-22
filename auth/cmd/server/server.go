package server

import (
	"auth/internal/core/ports"
	"auth/middlewares"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
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
	//app.Use(logger.New())
	logger, _ := zap.NewProduction()

	defer logger.Sync()
	v1 := app.Group("/v1")
	jwtMiddleware := middlewares.AuthMiddleware(string("SECRET:)"))

	userRoute := v1.Group("/user")
	userRoute.Post("/Register", s.userHandler.Register)
	userRoute.Post("/Login", s.userHandler.Login)
	userRoute.Get("/profile", jwtMiddleware, s.userHandler.Profile)

	err := app.Listen(":3000")
	if err != nil {
		logger.Error(err.Error())
	}
}
