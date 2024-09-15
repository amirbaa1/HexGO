package handler

import (
	"auth/internal/core/model"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func (h *UserHandl) Register(c *fiber.Ctx) error {
	auth := new(model.RegisterRequest)

	if err := c.BodyParser(auth); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid input data",
		})
	}

	err := h.app.Register(auth)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code": fiber.StatusInternalServerError,
			"msg":  err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"code":    fiber.StatusOK,
		"message": "user registered successfully",
	})
}

func (h *UserHandl) Login(c *fiber.Ctx) error {
	login := new(model.AuthRequest)
	if err := c.BodyParser(login); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":  "invalid input data",
			"status": fiber.StatusBadRequest,
		})
	}
	token, err := h.app.Login(login)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":  fiber.StatusInternalServerError,
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"code":    fiber.StatusOK,
		"message": "user logged in successfully",
		"token":   token,
	})
}

func (h *UserHandl) Profile(c *fiber.Ctx) error {
	userId := c.Locals("user").(*jwt.Token)

	profile, err := h.app.Profile(userId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code": fiber.StatusInternalServerError,
			"msg":  err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"code": fiber.StatusOK,
		"data": profile,
	})
}
