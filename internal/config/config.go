package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Loglevel string
	Port     string
	DBUrl    string
}

func PortInitialization() string {
	viper.SetDefault("PORT", "8080")
	return fmt.Sprintf(":%s", viper.GetString("PORT"))
}

func DBUrlInitialization() string {
	viper.SetDefault("DB_URL", "postgres://postgres:postgres@localhost:5432/messenger?sslmode=disable")
	return fmt.Sprintf(":%s", viper.GetString("PORT"))
}

func NewConfig() *Config {
	viper.AutomaticEnv()

	lvl := os.Getenv("LOG_LEVEL")
	if lvl == "" {
		lvl = "debug"
	}

	s := &Config{
		Loglevel: lvl,
		Port:     PortInitialization(),
		DBUrl:    DBUrlInitialization(),
	}

	return s
}
