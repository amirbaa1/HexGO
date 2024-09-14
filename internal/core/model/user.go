package model

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	Id         uuid.UUID `json:"id" gorm:"type:uuid;default:null"`
	Email      string    `json:"email" gorm:"column:email;type:varchar(255);not null"`
	Password   string    `json:"password" gorm:"column:password;type:varchar(255);not null"`
	CreateTime time.Time `json:"create_time" gorm:"autoCreateTime"`
}
