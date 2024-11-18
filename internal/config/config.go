package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/mfcbentes/go-oracle/internal/domain"
)

func LoadConfig() domain.Config {
	if err := godotenv.Load(); err != nil {
		log.Printf("Aviso: Não foi possível carregar .env: %v", err)
	}

	return domain.Config{
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBHost:     os.Getenv("DB_HOST"),
		DBService:  os.Getenv("DB_SERVICE"),
		DBPort:     os.Getenv("DB_PORT"),
	}
}
