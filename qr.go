package main

import (
	"github.com/skip2/go-qrcode"
)

// Функция для генерации QR-кода на основе UUID
func generateQRCode(uuid string) ([]byte, error) {
	qr, err := qrcode.New(uuid, qrcode.Medium)
	if err != nil {
		return nil, err
	}

	// Возвращаем изображение QR-кода
	img := qr.Image(256)
	return img, nil
}
