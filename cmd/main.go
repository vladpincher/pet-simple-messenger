package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	server "my-pet-simple-messenger/internal/api/http"
	"my-pet-simple-messenger/internal/config"
	"my-pet-simple-messenger/internal/database"
	"my-pet-simple-messenger/internal/logger"

	"go.uber.org/zap"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		exit := make(chan os.Signal, 1)
		signal.Notify(exit, os.Interrupt, syscall.SIGTERM)
		<-exit
		cancel()
	}()

	cfg := config.NewConfig()
	log := logger.NewLogger(cfg)

	db, err := database.InitDB(cfg.DBUrl)
	if err != nil {
		log.Fatal("Database error:", err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Fatal("Error closing database: %v", err)
		}
	}()

	s := server.NewServer(cfg, log, db)

	go func() {
		if err := s.Start(); err != nil {
			zap.Error(err)
		}
	}()

	<-ctx.Done()
	log.Info("Shutting down...")
	s.Close()
}
