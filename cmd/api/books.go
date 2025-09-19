// cmd/api/books.go

package main

import (
	"fmt"
	"go_skeleton/internal/data"
	"net/http"
	"time"
)

// Добавляем обработчик createBookHandler для эндпоинта `POST /v1/books`.
// Пока мы просто возвращаем текст.
func (app *application) createBookHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "create a new book")
}

// Добавляем обработчик showBookHandler для конечной точки `GET /v1/books/:id`.
// На данный момент мы получаем параметр id из URL-адреса и включаем его в ответ.
func (app *application) showBookHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	// Создаём новый экземпляр структуры Book, которая содержит идентификатор, взятый из
	// URL-адреса. Обратите внимание, мы намеренно не установили значение поля Year.
	book := data.Book{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "Effective Concurrency in Go",
		Author:    "Burak Serdar",
		Tags:      []string{"go", "programming", "concurrency"},
		Pages:     195,
	}

	// Сериализуем структуру в JSON и отправляем как HTTP-ответ.
	err = app.writeJSON(w, http.StatusOK, book, nil)
	if err != nil {
		app.logger.Println(err)
		http.Error(w, "The server could not process your request", http.StatusInternalServerError)
	}
}
