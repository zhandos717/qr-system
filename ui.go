package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// Функция для запуска приложения с UI
func runApp() {
	// Создание приложения
	myApp := app.New()
	myWindow := myApp.NewWindow("Получение билетов")

	// Поле для ввода номера
	entry := widget.NewEntry()
	entry.SetPlaceHolder("Введите ваш номер")

	// Лейбл для отображения результата
	responseLabel := widget.NewLabel("")

	// Пустое изображение для QR-кода
	qrImage := canvas.NewImageFromImage(nil)
	qrImage.FillMode = canvas.ImageFillContain

	// Кнопка для отправки данных и генерации QR-кода
	generateButton := widget.NewButton("Отправить", func() {
		number := entry.Text
		if number == "" {
			responseLabel.SetText("Пожалуйста, введите корректный номер.")
			return
		}

		// Отправка данных на сервер
		response, statusCode, err := sendDataToServer(number)
		if err != nil {
			responseLabel.SetText("Ошибка при отправке данных на сервер.")
			return
		}

		if statusCode == 404 {
			responseLabel.SetText("Пользователь не найден")
		} else {
			responseText := fmt.Sprintf("Пользователь найден:\nID: %d\nИмя: %s\nUUID: %s", response.Data.ID, response.Data.Name, response.Data.UUID)
			responseLabel.SetText(responseText)

			// Генерация QR-кода
			img, err := generateQRCode(response.Data.UUID)
			if err != nil {
				responseLabel.SetText("Ошибка при генерации QR-кода.")
				return
			}

			// Обновление изображения QR-кода
			qrImage.Image = img
			qrImage.Refresh()
		}
	})

	// Собираем интерфейс
	content := container.NewVBox(
		entry,
		generateButton,
		qrImage,
		responseLabel,
	)

	// Центрирование содержимого
	centeredContent := container.NewCenter(content)

	// Устанавливаем контент окна
	myWindow.SetContent(centeredContent)

	// Устанавливаем размеры окна
	myWindow.Resize(fyne.NewSize(400, 400))
	myWindow.CenterOnScreen()
	myWindow.ShowAndRun()
}
