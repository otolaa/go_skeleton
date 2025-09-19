// cmd/api/helpers.go

package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type envelope map[string]interface{}

// readIDParam получает параметр id из контекста запроса, преобразует его в
// целое число и возвращает его. Если операция не удалась,
// возвращает 0 и сообщение об ошибке.
func (app *application) readIDParam(r *http.Request) (int64, error) {
	param := chi.URLParam(r, "id")

	id, err := strconv.ParseInt(param, 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("invalid parameter")
	}

	return id, nil
}

// writeJSON вспомогательная функция для отправки ответов. Она принимает в качестве аргументов
// http.ResponseWriter, код ответа HTTP, данные для кодирования и заголовки, которые мы хотим
// включить в ответ.
func (app *application) writeJSON(w http.ResponseWriter, status int, data interface{}, headers http.Header) error {
	// Сериализуем данные в JSON. Возвращаем ошибку, если она возникает.
	js, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	// Добавим в наш JSON символ новой строки, чтобы было удобно его просматривать
	// в терминале.
	js = append(js, '\n')

	// На этом этапе мы знаем, что не столкнёмся с другими ошибками до отправки ответа,
	// поэтому можно безопасно добавлять любые заголовки.
	// Проходимся по мапе заголовков и добавляем каждый в http.ResponseWriter.
	for key, value := range headers {
		w.Header()[key] = value
	}

	// Добавим заголовок "Content-Type: application/json", затем укажем код состояния и
	// JSON-ответ.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}
