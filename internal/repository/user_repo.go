package repository

import (
	"my-pet-simple-messenger/internal/models"

	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user models.User) error {
	query := `
        INSERT INTO users (username, surname, email, phone_number, hash_password)
        VALUES (:username, :surname, :email, :phone_number, :hash_password)
    `
	_, err := r.db.NamedExec(query, user)
	return err
}

func (r *UserRepository) GetUserByUsername(username string) (*models.User, error) {
	var u models.User
	query := `SELECT * FROM users WHERE username = $1`
	err := r.db.Get(&u, query, username)
	if err != nil {
		return nil, err
	}
	return &u, nil
}
