package database

import (
	"database/sql"
	"fmt"
	"task-manager/pkg/logger"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	log := logger.Log

	dsn := fmt.Sprintf(
		"host=localhost port=5432 user=%s password=%s dbname=%s sslmode=disable",
		"admin", "password", "task_manager",
	)

	var err error
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Ошибка при подключении к БД: ", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal("БД недоступна: ", err)
	}

	log.Info("Подключено к PostgreSQL")
}
