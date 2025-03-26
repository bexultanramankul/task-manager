# Task Manager

Task Manager - это API-сервис для управления задачами, использующий Go, PostgreSQL и Docker.

## Установка зависимостей

```sh

go get github.com/lib/pq
go get github.com/spf13/viper
go get github.com/sirupsen/logrus

```

## Установка миграций

```sh

brew install golang-migrate

```

```sh

go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

```

## Запуск проекта

Перед запуском необходимо настроить конфигурацию.

### Конфигурация

Файл `config.yaml` должен находиться в `configs/` и содержать параметры подключения:

```yaml
server:
  port: "8080"

database:
  user: "postgres"
  password: "password"
  name: "task_manager"
  host: "localhost"
  port: "5432"
  sslmode: "disable"
```

### Запуск базы данных через Docker

```sh

docker-compose up -d

```

### Запуск приложения

```sh

go run cmd/server/main.go

```

## Структура проекта

```plaintext
├── cmd
│   ├── server
│   │   └── main.go       # Точка входа в приложение
│
├── internal
│   ├── config           # Загрузка конфигурации
│   ├── database         # Инициализация базы данных
│   ├── models           # Определение моделей данных
│
├── pkg
│   ├── logger           # Настройка логирования
│
├── db
│   ├── migrations       # SQL-миграции базы данных
│
├── configs
│   └── config.yaml      # Файл конфигурации
│
├── deployments          # Файлы деплоя (Docker, Kubernetes и т. д.)
│
├── Makefile             # Скрипты сборки и запуска
├── docker-compose.yml   # Конфигурация Docker
```

## API

### Проверка работы сервера

```sh

curl http://localhost:8080/

```

Ответ:

```json
{"message": "Task Manager API is running!"}
```

## Завершение работы

Для остановки сервера используйте команду:

```sh

CTRL + C

```

Для остановки Docker-контейнера:

```sh

docker-compose down

```