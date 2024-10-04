package config

import (
	"github.com/joho/godotenv"
	"log"
)

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Erro ao carregar o arquivo .env: %v", err)
	}
}
