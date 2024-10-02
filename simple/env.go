package simple

import (
	"log"

	"github.com/joho/godotenv"
)

// Функция для загрузки .env файла
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Ошибка загрузки .env файла")
	}
}
