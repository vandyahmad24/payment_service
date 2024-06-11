package config

import (
	"log"

	"github.com/vandyahmad24/alat-bantu/config"
)

var (
	configPath = "./.env"
	EnvConfig  = NewConfig()
)

type Config struct {
	PORT                string
	JWT_SECRET          string
	TRANSACTION_SERVICE string
	CUSTOMER_SERVICE    string
}

func NewConfig() *Config {
	if err := config.LoadEnv(configPath); err != nil {
		log.Fatalf("Application dimissed. Application cannot find %s to run this application", err.Error())
	}

	return &Config{
		PORT:                config.GetEnv("PORT", ""),
		JWT_SECRET:          config.GetEnv("JWT_SECRET", ""),
		TRANSACTION_SERVICE: config.GetEnv("TRANSACTION_SERVICE", ""),
		CUSTOMER_SERVICE:    config.GetEnv("CUSTOMER_SERVICE", ""),
	}

}
