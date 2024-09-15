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

func (h *Handler) UpdateBook(c *fiber.Ctx) error {
	reqBook := new(model.RequestBookUpdate)
	id := c.Params("id")

	if err := c.BodyParser(reqBook); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":  err,
			"status": fiber.StatusBadRequest,
		})
	}

	result, err := h.service.UpdateBook(reqBook, id)
	if err != nil {
		if err.Error() == "book not found" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error":  "Book not found",
				"status": fiber.StatusNotFound,
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":  err.Error(),
			"status": fiber.StatusInternalServerError,
		})
	}
	//if result.Error != nil {
	//	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
	//		"error":  result.Error,
	//		"status": fiber.StatusBadRequest,
	//	})
	//}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "Book updated",
		"data":    result,
	})

}

func (h *Handler) GetBookById(c *fiber.Ctx) error {
	id := c.Params("id")
	result, err := h.service.GetBookById(id)

	if err != nil {
		if err.Error() == "book not found" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error":  "Book not found",
				"status": fiber.StatusNotFound,
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":  err.Error(),
			"status": fiber.StatusInternalServerError,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": fiber.StatusOK,
		"data":   result,
	})

}
