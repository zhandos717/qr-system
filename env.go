package main

import (
	"log"

	"github.com/joho/godotenv"
)

// Функция для загрузки .env файла
func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Ошибка загрузки .env файла")
	}
}
