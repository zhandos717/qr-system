package main

import (
	"github.com/zhandos717/qr-system/simple"
)

func main() {
	// Загрузка переменных окружения
	simple.LoadEnv()

	// Запуск приложения
	simple.RunApp()
}
