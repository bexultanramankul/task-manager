package main

import (
	"fmt"
	"net/http"
	"os"
	"task-manager/pkg/config"
	"task-manager/pkg/database"
	"task-manager/pkg/logger"
)

func main() {
	logger.InitLogger()
	log := logger.Log

	log.Info("Загрузка конфигурации...")
	config.LoadConfig()

	log.Info("Подключение к базе данных...")
	database.InitDB()

	port := ":" + os.Getenv("SERVER_PORT")
	log.Infof("Сервер запущен на http://localhost%s", port)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Task Manager API is running!")
	})

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal("Ошибка запуска сервера:", err)
	}
}
