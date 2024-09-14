package ports

import "HexGO/internal/core/model"

type UserService interface {
	Register(register *model.RegisterRequest) error
	Login(login *model.AuthRequest) (model.AuthResponse, error)
}
