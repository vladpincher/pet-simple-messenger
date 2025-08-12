package main

import (
	"fmt"

	"my-pet-simple-messenger/internal/api/database"
	server "my-pet-simple-messenger/internal/api/http"
	"my-pet-simple-messenger/internal/config"
	"my-pet-simple-messenger/internal/logger"
)

func main() {

	cfg := config.NewConfig()
	logg := logger.NewLogger(cfg)

	db, err := database.InitDB(cfg.DBUrl)
	if err != nil {
		fmt.Println("database error:", err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			fmt.Printf("Ошибка при закрытии БД: %v\n", err)
		}
	}()

	s := server.NewServer(cfg, logg, db)
	defer s.Close()

	if err := s.Start(); err != nil {
		fmt.Println("server error:", err)
	}
}
