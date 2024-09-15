package repository

import (
	"book/internal/core/ports"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

var _ ports.BookRepository = (*Repository)(nil)

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}
