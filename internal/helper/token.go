package helper

import (
	"HexGO/internal/core/model"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var Secret = []byte("SECRET:)")

func GenerateToken(user *model.User) (string, error) {
	claims := jwt.MapClaims{
		"user-id": user.Id,
		"email":   user.Email,
		"exp":     time.Now().Add(time.Minute * 15).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(Secret)
}
