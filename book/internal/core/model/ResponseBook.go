package model

import (
	"github.com/google/uuid"
	"time"
)

type ResponseBook struct {
	Id        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	AuthorId  uuid.UUID `json:"authorId"`
	Author    Author    `json:"author"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
