package models

import "time"

type User struct {
	Id           int       `db:"id" json:"id"`
	Username     string    `db:"username" json:"username"`
	Surname      string    `db:"surname" json:"surname"`
	Email        string    `db:"email" json:"email"`
	Phone        string    `db:"phone_number" json:"phone_number"`
	HashPassword string    `db:"hash_password" json:"hash_password"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
}
