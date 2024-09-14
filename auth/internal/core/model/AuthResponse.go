package model

import "errors"

type AuthResponse struct {
	Token string `json:"token"`
}

func (a *AuthResponse) Validate() error {
	if a.Token == "" {
		return errors.New("token is empty")
	}
	return nil
}
