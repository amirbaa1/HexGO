package handler

import "book/internal/core/ports"

type Handler struct {
	service ports.BookService
}

var _ ports.BookHandler = (*Handler)(nil)

func NewHandler(service ports.BookService) *Handler {
	return &Handler{
		service: service,
	}
}
