package config

import (
	"os"

	"github.com/joho/godotenv"
)

func Config(key string) string {
	err := godotenv.Load(".env")
	if err != nil{
		return "Gagal memuat file env"
	}
	return os.Getenv(key)
}