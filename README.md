# Go step by step 🍔

```
go mod init
go mod tidy
go run .
```

## Skeleton ☠️

```
.
|---bin
|---cmd
|   +---api
|       +---main.go
|---internal
|---migrations
|---remote
|---go.mod
|---Makefile
```

## Info for folder catalog ⚔️

```
bin — скомпилированные двоичные файлы, готовые к развёртыванию на рабочем сервере;

cmd/api —  основная функция приложения;

internal — различные вспомогательные пакеты;

migrations — файлы миграции для базы данных;

remote — файлы конфигурации и сценарии настроек для производственного сервера;

go.mod — информация о зависимостях проекта, версиях и путях к модулям;

Makefile — инструкции по автоматизации частых административных задач — проверка кода, создание двоичных файлов и выполнения миграций.
```

## Run 🚀

```
killall -9 go
$ go run ./cmd/api
```

## point API and REST

| Метод | Эндпоинт | Хендлер | Действие |
| :---:   | :---: | :---: | :---: |
| GET | /v1/healthcheck   | healthcheckHandler | Информация о приложении |
| GET | /v1/books|  | listBooksHandler | список книг |
| POST | /v1/books   | createHandler | Создает новую книгу |
| GET | /v1/books/{id}   | showBookHandler | Детали книги |
| PUT | /v1/books/{id}   | editBookHandler | Обновляет информацию |
| DELETE | /v1/books/{id}   | deleteBookHandler | Удаляет книгу |

## use curl in terminal

```
$ curl -i localhost:8000/v1/healthcheck
status: available
environment: development
version: 1.0.0

$ curl -i -X POST localhost:8000/v1/books
create a new book

$ curl -i localhost:8000/v1/books/12
show the details of book 12
```