package service

import (
	"context"
	"errors"
	"fmt"

	"my-pet-simple-messenger/internal/models"
	"my-pet-simple-messenger/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrUserAlreadyExists = errors.New("user already exists")
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(ctx context.Context, user models.User) error {
	// 1. Проверка — есть ли уже такой пользователь
	_, err := s.repo.GetUserByUsername(ctx, user.Username)
	if err == nil {
		return ErrUserAlreadyExists
	}

	// 2. Хешируем пароль
	hash, err := bcrypt.GenerateFromPassword([]byte(user.HashPassword), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("hash password: %w", err)
	}
	user.HashPassword = string(hash)

	if err := s.repo.CreateUser(ctx, user); err != nil {
		return fmt.Errorf("create user: %w", err)
	}

	return nil
}
