package handler

import (
	"book/internal/core/model"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) CreateBook(c *fiber.Ctx) error {
	bookNew := new(model.RequestBook)
	if err := c.BodyParser(bookNew); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":  err,
			"status": fiber.StatusBadRequest,
		})
	}
	err := h.service.CreateServiceBook(bookNew)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":  err,
			"status": fiber.StatusBadRequest,
		})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  fiber.StatusCreated,
		"message": "Book created",
	})
}

func (h *Handler) GetAllBooks(c *fiber.Ctx) error {
	getAll, err := h.service.GetAllBooks()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":  err,
			"status": fiber.StatusBadRequest,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": fiber.StatusOK,
		"data":   getAll,
	})
}
