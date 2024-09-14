package model

import "errors"

type AuthRequest struct {
	Email                string `json:"email"`
	Password             string `json:"password"`
}

func (a *AuthRequest) Validate() error {
	if a.Email == "" {
		return errors.New("Email can't be empty")
	}
	if a.Password == "" {
		return errors.New("Password can't be empty")
	}
	return nil
}
