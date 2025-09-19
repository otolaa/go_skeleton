// internal/data/books.go
package data

import (
	"time"
)

type Book struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"` // "-" можно использовать, чтобы поле НЕ отображалось в JSON.
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	Year      int32     `json:"year,omitempty"` // "omitempty" скрывает поле, если значение поля пусто
	Tags      []string  `json:"tags,omitempty"` // "omitempty" скрывает поле, если значение поля пусто
	Pages     Pages     `json:"pages"`
}
