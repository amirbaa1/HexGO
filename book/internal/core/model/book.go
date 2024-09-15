package model

import (
	"github.com/google/uuid"
	"time"
)

type Book struct {
	Id         uuid.UUID `json:"id" gorm:"type:uuid;default:null"`
	Title      string    `json:"title"`
	AuthorId   uuid.UUID `json:"author_id"`
	Author     Author    `json:"author"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}

type Author struct {
	Id        uuid.UUID `json:"id" gorm:"type:uuid;default:null"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
}
