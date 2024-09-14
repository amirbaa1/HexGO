package handler

import "HexGO/internal/core/ports"

type UserHandl struct {
	app ports.UserService
}

var _ ports.UserHandler = (*UserHandl)(nil)

func NewUserHandler(app ports.UserService) *UserHandl {
	return &UserHandl{app: app}
}
