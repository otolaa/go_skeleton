package data

import (
	"fmt"
	"strconv"
)

// Объявляем пользовательский тип Pages, который имеет базовый тип int32
// такой же, как был у поля Pages структуры Book
type Pages int32

// Реализуем метод MarshalJSON(), чтобы тип Pages удовлетворял интерфейсу json.Marshaler
// Метод должен возвращать количество страниц в книге в кодировке JSON в таком виде:
// "<количество> страниц"
func (p Pages) MarshalJSON() ([]byte, error) {
	// Создаем строку, которая содержит количество страниц в нужном нам формате
	jsonValue := fmt.Sprintf("%d pages", p)

	// Используем функцию strconv.Quote() для созданной строки, чтобы заключить её в
	// двойные кавычки. Это нужно, так как строки в JSON заключены в кавычки.
	quotedJSONValue := strconv.Quote(jsonValue)

	// Преобразовываем строку в массив байтов и возвращаем её
	return []byte(quotedJSONValue), nil
}
