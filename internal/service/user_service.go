package service

import (
	"errors"
	"my-pet-simple-messenger/internal/models"
	"my-pet-simple-messenger/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Register(username, surname, email, phone, password string) error {
	// 1. Проверка — есть ли уже такой пользователь
	_, err := s.repo.GetUserByUsername(username)
	if err == nil {
		return errors.New("user already exists")
	}

	// 2. Хешируем пароль
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// 3. Создаем пользователя
	user := models.User{
		Username:     username,
		Surname:      surname,
		Email:        email,
		Phone:        phone,
		HashPassword: string(hash),
	}
	return s.repo.CreateUser(user)
}
