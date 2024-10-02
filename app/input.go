package app

import (
	"fyne.io/fyne/v2/widget"
)

// InputField представляет собой компонент поля ввода
type InputField struct {
	Entry *widget.Entry
}

// NewInputField создаёт новый компонент поля ввода
func NewInputField(placeholder string) *InputField {
	entry := widget.NewEntry()
	entry.SetPlaceHolder(placeholder)

	return &InputField{Entry: entry}
}

// GetText возвращает текст, введённый в поле
func (f *InputField) GetText() string {
	return f.Entry.Text
}
