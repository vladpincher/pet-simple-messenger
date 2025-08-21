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
	viper.SetDefault("DB_PROTOCOL", "postgres")
	viper.SetDefault("DB_USER", "postgres")
	viper.SetDefault("DB_PASSWORD", "postgres")
	viper.SetDefault("DB_HOST", "localhost")
	viper.SetDefault("DB_PORT", "5432")
	viper.SetDefault("DB_NAME", "messenger")
	viper.SetDefault("DB_SSLMODE", "disable")

	return fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=%s",
		viper.GetString("DB_PROTOCOL"),
		viper.GetString("DB_USER"), viper.GetString("DB_PASSWORD"),
		viper.GetString("DB_HOST"), viper.GetString("DB_PORT"),
		viper.GetString("DB_NAME"),
		viper.GetString("DB_SSLMODE"))
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
