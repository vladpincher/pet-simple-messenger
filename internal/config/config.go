package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Loglevel string
	Port     string
}

func PortInitialization() string {
	viper.AutomaticEnv()
	viper.SetDefault("PORT", "8080")
	return fmt.Sprintf(":%s", viper.GetString("PORT"))
}

func NewConfig() *Config {

	lvl := os.Getenv("LOG_LEVEL")
	if lvl == "" {
		lvl = "debug"
	}

	s := &Config{
		Loglevel: lvl,
		Port:     PortInitialization(),
	}

	return s
}
