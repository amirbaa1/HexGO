package server

import (
	"auth/internal/core/ports"
	"auth/middlewares"
	"log"

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

func StartRabbitMQ() (ports.MessagingPort, error) {
	// conn, err := ConnectRabbit()

	// if err != nil {
	// 	log.Println("Failed to connect to RabbitMQ", zap.Error(err))
	// }
	// defer conn.Close()

	// rabbitMQClient, err := NewRabbitMQ(conn)
	// if err != nil {
	// 	log.Println("Failed to create RabbitMQ client: %v", zap.Error(err))
	// }
	// defer rabbitMQClient.Close()

	// log.Print("Connected to RabbitMQ")

	// err = rabbitMQClient.CreateQueueDeclare("emailQueue", true, false)
	// if err != nil {
	// 	log.Print("Failed to create queue: %v", zap.Error(err))
	// }

	// select {}
	conn, err := ConnectRabbit()
	if err != nil {
		log.Println("Failed to connect to RabbitMQ", err)
		return nil, err
	}

	rabbitMQClient, err := NewRabbitMQ(conn)
	if err != nil {
		log.Printf("Failed to create RabbitMQ client: %v", err)
		return nil, err
	}

	log.Print("Connected to RabbitMQ")

	err = rabbitMQClient.CreateQueueDeclare("emailQueue", true, false)
	if err != nil {
		log.Printf("Failed to create queue: %v", zap.Error(err))
		return nil, err
	}

	return rabbitMQClient, nil
}
