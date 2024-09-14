package ports

import "HexGO/internal/core/model"

type UserRepository interface {
	Register(auth *model.User) error
	Login(login *model.AuthRequest) (*model.User, error)
}
