package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Internal() {
	app := fiber.New()
	app.Use(logger.New())

	err := app.Listen(":3003")
	if err != nil {
		log.Fatal(err)
	}
}
