package repository

import (
	"HexGO/internal/core/ports"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

var _ ports.UserRepository = (*UserRepository)(nil)

func NewRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}
