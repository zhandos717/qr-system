package app

import (
	"log"
	"time"

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

	// Создаём новый компонент поля ввода
	inputField := NewInputField("Введите ваш номер")

	// Лейбл для отображения результата
	responseLabel := widget.NewLabel("")

	// Пустое изображение для QR-кода
	qrImage := canvas.NewImageFromImage(nil)
	qrImage.FillMode = canvas.ImageFillOriginal

	var countdownTimer *time.Timer // Переменная для хранения таймера

	// Функция для генерации QR-кода на основе UUID пользователя и отправки данных на сервер
	generateQR := func() {
		number := inputField.GetText()
		if number == "" {
			responseLabel.Importance = widget.DangerImportance
			responseLabel.SetText("Пожалуйста, введите корректный номер.")
			qrImage.Image = nil
			qrImage.Refresh()
			return
		}

		// Отправка данных на сервер и обработка ответа
		response, statusCode, err := sendDataToServer(number)
		if err != nil {
			responseLabel.SetText("Ошибка при отправке данных на сервер.")
			log.Println(err)
		} else {
			if statusCode == 404 {
				responseLabel.Importance = widget.DangerImportance
				responseLabel.SetText("Пользователь не найден")
				qrImage.Image = nil // Очищаем изображение при неудаче
				qrImage.Refresh()
			} else {

				responseLabel.Importance = widget.SuccessImportance

				responseLabel.SetText("Пользователь найден")

				// Генерация QR-кода на основе UUID
				img, err := generateQRCode(response.Data.UUID)
				if err != nil {
					qrImage.Image = nil
					qrImage.Refresh()
					responseLabel.SetText("Ошибка при генерации QR-кода.")
					return
				}

				// Обновление изображения QR-кода
				qrImage.Image = img
				qrImage.Refresh() // Обязательно обновляем изображение

				// Если таймер уже запущен, останавливаем его
				if countdownTimer != nil {
					countdownTimer.Stop()
				}

				// Запускаем новый обратный отсчет
				countdownTimer = time.AfterFunc(10*time.Second, func() {
					resetUI(inputField, responseLabel, qrImage)
				})
			}
		}
	}

	// Кнопка для генерации QR-кода и отправки данных
	generateButton := widget.NewButton("Отправить", func() {
		generateQR()
	})

	// Настройка размера кнопки
	generateButton.Resize(fyne.NewSize(200, 100))     // Ширина 200 пикселей, высота 50 пикселей
	generateButton.Importance = widget.HighImportance // Устанавливаем кнопку как важную

	// Добавление цветового фона для кнопки
	//	generateButton.Style.BackgroundColor = fyne.NewColor(0, 0, 255) // Синий цвет
	//	generateButton.Style.Color = fyne.NewColor(255, 255, 255)       // Белый цвет текста

	bottomContainer := container.NewVBox(

		inputField.Entry,
		generateButton,
	)

	// Собираем интерфейс
	content := container.NewVBox(
		qrImage,
		container.NewCenter(responseLabel),
		bottomContainer,
	)

	// Устанавливаем контент окна
	myWindow.SetContent(content)

	// Устанавливаем размеры окна
	myWindow.Resize(fyne.NewSize(400, 400))
	myWindow.CenterOnScreen()
	myWindow.ShowAndRun()
}

// Функция для сброса UI к начальному состоянию
func resetUI(inputField *InputField, responseLabel *widget.Label, qrImage *canvas.Image) {
	inputField.Entry.SetText("") // Сбрасываем текст в поле ввода
	responseLabel.SetText("")    // Очищаем текст результата
	qrImage.Image = nil          // Очищаем изображение QR-кода
	qrImage.Refresh()            // Обновляем изображение
}
