package ports

import "auth/internal/core/model"

type UserRepository interface {
	Register(auth *model.User) error
	//Login(login *model.AuthRequest) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
}
