package main

import (
	"fmt"
	"net/http"
	"os"
	"task-manager/pkg/config"
	"task-manager/pkg/database"
	"task-manager/pkg/logger"
)

func handleRoot(w http.ResponseWriter, r *http.Request) {
	if _, err := fmt.Fprintln(w, "Task Manager API is running!"); err != nil {
		logger.Log.Error("Ошибка записи в ResponseWriter:", err)
	}
}

func startServer() {
	port := ":" + os.Getenv("SERVER_PORT")
	logger.Log.Infof("Сервер запущен на http://localhost%s", port)

	http.HandleFunc("/", handleRoot)

	if err := http.ListenAndServe(port, nil); err != nil {
		logger.Log.Fatal("Ошибка запуска сервера:", err)
	}
}

func main() {
	logger.InitLogger()
	logger.Log.Info("Загрузка конфигурации...")
	config.LoadConfig()

	logger.Log.Info("Подключение к базе данных...")
	database.InitDB()

	startServer()
}
