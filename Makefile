# Настройки базы данных
DB_URL="postgres://postgres:password@localhost:5432/task_manager?sslmode=disable"
MIGRATIONS_DIR=db/migrations
NAME=
VERSION=1

.PHONY: up down force goto version create

# Применить все миграции
up:
	migrate -path $(MIGRATIONS_DIR) -database $(DB_URL) up

# Откатить последнюю миграцию
down:
	migrate -path $(MIGRATIONS_DIR) -database $(DB_URL) down 1

# Откатить ВСЕ миграции
down-all:
	migrate -path $(MIGRATIONS_DIR) -database $(DB_URL) down

# Принудительно задать версию миграции (указать VERSION=номер)
force:
	migrate -path $(MIGRATIONS_DIR) -database $(DB_URL) force $(VERSION)

# Перейти к конкретной версии миграции (указать VERSION=номер)
goto:
	migrate -path $(MIGRATIONS_DIR) -database $(DB_URL) goto $(VERSION)

# Посмотреть текущую версию миграций
version:
	migrate -path $(MIGRATIONS_DIR) -database $(DB_URL) version

# Создать новую миграцию (указать NAME=имя_миграции)
create:
	migrate create -ext sql -dir $(MIGRATIONS_DIR) -seq $(NAME)

# Используй migrate force, чтобы сбросить dirty-флаг
migrate-force:
	migrate -path db/migrations -database $(DB_URL) force $(VERSION)