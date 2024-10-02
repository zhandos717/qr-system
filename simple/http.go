package simple

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
)

// Структура для ответа от сервера
type ResponseData struct {
	Data struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		UUID string `json:"uuid"`
	} `json:"data"`
}

// Функция для отправки данных на сервер и обработки ответа
func sendDataToServer(number string) (*ResponseData, int, error) {
	apiURL := os.Getenv("API_URL")
	bearerToken := os.Getenv("BEARER_TOKEN")

	// Подготовка JSON данных
	requestData := map[string]string{
		"number": number,
	}
	jsonData, err := json.Marshal(requestData)
	if err != nil {
		return nil, 0, err
	}

	// Создание HTTP-запроса
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, 0, err
	}

	// Добавление заголовков
	req.Header.Set("Authorization", "Bearer "+bearerToken)
	req.Header.Set("Content-Type", "application/json")

	// Выполнение запроса
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()

	// Чтение тела ответа с использованием io
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, 0, err
	}

	// Если код ответа 404, возвращаем пустые данные и код
	if resp.StatusCode == 404 {
		return nil, 404, nil
	}

	// Парсинг JSON-ответа
	var responseData ResponseData
	err = json.Unmarshal(body, &responseData)
	if err != nil {
		return nil, resp.StatusCode, err
	}

	return &responseData, resp.StatusCode, nil
}
