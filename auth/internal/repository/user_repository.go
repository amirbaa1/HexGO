package repository

import (
	"HexGO/internal/core/model"
	"errors"
)

func (r *UserRepository) Register(user *model.User) error {
	return r.db.Create(&user).Error
}

func (r *UserRepository) Login(login *model.AuthRequest) (*model.User, error) {
	var user model.User

	err := r.db.Where("email = ?", login.Email).First(&user).Error
	if err != nil {
		return nil, errors.New("user not found")
	}

	return &user, nil
}
