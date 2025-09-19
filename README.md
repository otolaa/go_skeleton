# Go step by step 🍔

```
go mod init
go mod tidy
go run .
```

## skeleton

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

## info for folder catalog

```
bin — скомпилированные двоичные файлы, готовые к развёртыванию на рабочем сервере;

cmd/api —  основная функция приложения;

internal — различные вспомогательные пакеты;

migrations — файлы миграции для базы данных;

remote — файлы конфигурации и сценарии настроек для производственного сервера;

go.mod — информация о зависимостях проекта, версиях и путях к модулям;

Makefile — инструкции по автоматизации частых административных задач — проверка кода, создание двоичных файлов и выполнения миграций.
```


## Run

```
killall -9 go
$ go run ./cmd/api
```

## use curl in terminal, the -i curl returm header and body server:

```
$ curl -i localhost:8000/v1/healthcheck
```

## point API and REST

| Метод | Эндпоинт | Хендлер | Действие |
| :---:   | :---: | :---: | :---: |
| GET | /v1/healthcheck   | healthcheckHandler | Информация о приложении |
| POST | /v1/books   | createHandler | Создает новую книгу |
| GET | /v1/books/{id}   | showBookHandler | Детали книги |


## test curl

```
$ curl localhost:4000/v1/healthcheck
status: available
environment: development
version: 1.0.0
```

```
$ curl -X POST localhost:4000/v1/books
create a new book
```

```
$ curl localhost:4000/v1/books/12
show the details of book 12
```