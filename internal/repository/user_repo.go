package repository

import (
	"context"
	"my-pet-simple-messenger/internal/models"

	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(ctx context.Context, user models.User) error {
	query := `
        INSERT INTO users (username, surname, email, phone_number, hash_password, created_at)
        VALUES (:username, :surname, :email, :phone_number, :hash_password, :created_at)
    `
	_, err := r.db.NamedExec(query, user)
	return err
}

func (r *UserRepository) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	var u models.User
	query := `
		SELECT 
			id,
			username,
			surname,
			email,
			phone_number,
			hash_password,
			created_at
		FROM users 
		WHERE username = $1
	`
	err := r.db.Get(&u, query, username)
	if err != nil {
		return nil, err
	}
	return &u, nil
}
