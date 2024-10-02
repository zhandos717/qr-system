package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/skip2/go-qrcode"
)

func main() {
	// Создание приложения
	myApp := app.New()
	myWindow := myApp.NewWindow("QR Code Generator")

	// Поле для ввода номера
	entry := widget.NewEntry()
	entry.SetPlaceHolder("Введите ваш номер")

	// Лейбл для отображения ошибок
	errorLabel := widget.NewLabel("")

	// Пустое изображение для QR-кода
	qrImage := canvas.NewImageFromImage(nil)
	qrImage.FillMode = canvas.ImageFillContain

	// Функция для генерации QR-кода
	generateQR := func(number string) {
		if number == "" {
			errorLabel.SetText("Please enter a valid number.")
			return
		}
		// Генерация QR-кода
		qr, err := qrcode.New(number, qrcode.Medium)
		if err != nil {
			errorLabel.SetText("Failed to generate QR code.")
			return
		}

		img := qr.Image(256)
		qrImage.Image = img
		qrImage.Refresh()
		errorLabel.SetText("")
	}

	// Кнопка для генерации QR-кода
	generateButton := widget.NewButton("Generate QR Code", func() {
		generateQR(entry.Text)
	})

	// Собираем интерфейс
	content := container.NewVBox(
		entry,
		generateButton,
		qrImage,
		errorLabel,
	)

	// Устанавливаем контент окна
	myWindow.SetContent(content)

	// Запуск приложения
	myWindow.Resize(fyne.NewSize(400, 400))
	myWindow.ShowAndRun()
}
