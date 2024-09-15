package model

import "github.com/google/uuid"

type ProfileResponse struct {
	Id    uuid.UUID `json:"id"`
	Email string    `json:"email"`
	Token string    `json:"token"`
}
