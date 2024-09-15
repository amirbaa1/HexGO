package service

import (
	"book/internal/core/ports"
)

type Service struct {
	bookRepository ports.BookRepository
}

var _ ports.BookService = (*Service)(nil)

func NewServiceBook(bookRepository ports.BookRepository) *Service {
	return &Service{
		bookRepository: bookRepository,
	}
}
