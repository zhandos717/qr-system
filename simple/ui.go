package simple

import (
	"fmt"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// Функция для запуска приложения с UI
func RunApp() {
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

	// Функция для генерации QR-кода на основе UUID пользователя и отправки данных на сервер
	generateQR := func(number string) {
		if number == "" {
			responseLabel.SetText("Пожалуйста, введите корректный номер.")
			return
		}

		// Отправка данных на сервер и обработка ответа
		response, statusCode, err := sendDataToServer(number)
		if err != nil {
			responseLabel.SetText("Ошибка при отправке данных на сервер.")
			log.Println(err)
		} else {
			if statusCode == 404 {
				responseLabel.SetText("Пользователь не найден")
			} else {
				responseText := fmt.Sprintf("Пользователь найден:\nID: %d\nИмя: %s\nUUID: %s", response.Data.ID, response.Data.Name, response.Data.UUID)
				responseLabel.SetText(responseText)

				// Генерация QR-кода на основе UUID
				img, err := generateQRCode(response.Data.UUID)
				if err != nil {
					responseLabel.SetText("Ошибка при генерации QR-кода.")
					return
				}

				// Обновление изображения QR-кода
				qrImage.Image = img
				qrImage.Refresh()
			}
		}
	}

	// Кнопка для генерации QR-кода и отправки данных
	generateButton := widget.NewButton("Отправить", func() {
		generateQR(entry.Text)
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
