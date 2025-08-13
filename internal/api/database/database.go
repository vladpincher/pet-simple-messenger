package database

import (
	"fmt"
	"time"

	retry "github.com/avast/retry-go/v4"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func InitDB(connStr string) (*sqlx.DB, error) {

	var db *sqlx.DB

	err := retry.Do(
		func() error {
			var err error
			db, err = sqlx.Open("postgres", connStr)
			if err != nil {
				return fmt.Errorf("ошибка при подключении к БД: %w", err)
			}

			if err := db.Ping(); err != nil {
				return fmt.Errorf("БД не отвечает: %w", err)
			}

			return nil
		},
		retry.Attempts(5),
		retry.Delay(1*time.Second),
		retry.DelayType(retry.FixedDelay),
		retry.LastErrorOnly(true),
	)

	if err != nil {
		return nil, err
	}

	return db, nil
}
