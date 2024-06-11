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
	PORT                   string
	MYSQL_USERNAME         string
	MYSQL_PASSWORD         string
	MYSQL_HOST             string
	MYSQL_PORT             string
	MYSQL_DATABASE         string
	PAYMENT_METHOD_SERVICE string
}

func NewConfig() *Config {
	if err := config.LoadEnv(configPath); err != nil {
		log.Fatalf("Application dimissed. Application cannot find %s to run this application", err.Error())
	}

	return &Config{
		PORT:                   config.GetEnv("PORT", ""),
		MYSQL_USERNAME:         config.GetEnv("MYSQL_USERNAME", ""),
		MYSQL_PASSWORD:         config.GetEnv("MYSQL_PASSWORD", ""),
		MYSQL_HOST:             config.GetEnv("MYSQL_HOST", ""),
		MYSQL_PORT:             config.GetEnv("MYSQL_PORT", ""),
		MYSQL_DATABASE:         config.GetEnv("MYSQL_DATABASE", ""),
		PAYMENT_METHOD_SERVICE: config.GetEnv("PAYMENT_METHOD_SERVICE", ""),
	}

}
