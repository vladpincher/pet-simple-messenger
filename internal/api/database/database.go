package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func InitDB(connStr string) (*sql.DB, error) {

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("ошибка при подключении к БД: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("БД не отвечает: %w", err)
	}

	return db, nil
}
