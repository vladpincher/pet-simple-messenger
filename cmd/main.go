package main

import (
	"fmt"

	"my-pet-simple-messenger/http/server"
	"my-pet-simple-messenger/internal/config"
	"my-pet-simple-messenger/internal/logger"
)

func main() {

	cfg := config.NewConfig()
	logg := logger.NewLogger(cfg)

	s := server.NewServer(cfg, logg)
	defer s.Close()

	if err := s.Start(); err != nil {
		fmt.Println("server error:", err)
	}
}
