package main

import (
	 "github.com/zhandos717/qr-system/app"
)

func main() {
	// Загрузка переменных окружения
	app.LoadEnv()

	// Запуск приложения
	app.RunApp()
}
