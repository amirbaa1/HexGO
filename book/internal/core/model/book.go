package model

import (
	"github.com/google/uuid"
	"time"
)

type Book struct {
	Id         uuid.UUID `json:"id"`
	Title      string    `json:"title"`
	Author     Author    `json:"author"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}

type Author struct {
	Id        uuid.UUID `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
}
