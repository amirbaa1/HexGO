package service

import (
	"auth/internal/core/model"
	"auth/internal/helper"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"strings"
	"time"
)

func (s *UserService) Register(register *model.RegisterRequest) error {
	if register == nil {
		return errors.New("register request cannot be nil")
	}
	if strings.TrimSpace(register.Email) == "" {
		return errors.New("email cannot be empty")
	}

	if register.Password != register.PasswordConfirmation {
		return errors.New("the passwords are not equal")
	}

	var err error
	existingUser, err := s.userRepository.FindByEmail(register.Email)
	if err == nil && existingUser != nil {
		return errors.New("user already exists")
	}

	hashPass, err := helper.HashPassword(register.Password)
	if err != nil {
		return err
	}
	userNew := &model.User{
		Id:         uuid.New(),
		Email:      register.Email,
		Password:   hashPass,
		CreateTime: time.Now(),
	}

	err = s.userRepository.Register(userNew)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) Login(login *model.AuthRequest) (model.AuthResponse, error) {
	if login == nil {
		return model.AuthResponse{}, errors.New("login email cannot be nil")
	}
	if err := login.Validate(); err != nil {
		return model.AuthResponse{}, err
	}

	user, err := s.userRepository.FindByEmail(login.Email)
	if err != nil {
		return model.AuthResponse{}, err
	}
	errPass := helper.CheckPassword(user.Password, login.Password)
	if errPass != nil {
		return model.AuthResponse{}, errors.New("invalid email or password")
	}
	token, err := helper.GenerateToken(user)

	return model.AuthResponse{
		token,
	}, nil
}

func (s *UserService) Profile(token *jwt.Token) (model.ProfileResponse, error) {
	claims := token.Claims.(jwt.MapClaims)
	email := claims["email"].(string)

	userGet, err := s.userRepository.FindByEmail(email)
	if err != nil {
		return model.ProfileResponse{}, err
	}

	user := &model.ProfileResponse{
		Id:    userGet.Id,
		Email: userGet.Email,
		//Token: string(claims),
	}
	return *user, nil
}
