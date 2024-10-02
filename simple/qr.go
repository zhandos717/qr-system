package simple

import (
	"image"

	"github.com/skip2/go-qrcode"
)

// Функция для генерации QR-кода на основе UUID
func generateQRCode(uuid string) (image.Image, error) {
	qr, err := qrcode.New(uuid, qrcode.Medium)
	if err != nil {
		return nil, err
	}

	// Возвращаем изображение QR-кода
	return qr.Image(256), nil
}
