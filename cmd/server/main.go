package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"task-manager/internal/database"
	"time"

	"task-manager/internal/config"
	"task-manager/pkg/logger"
)

// Обработчик корневого маршрута
func handleRoot(w http.ResponseWriter, r *http.Request) {
	if _, err := fmt.Fprintln(w, "Task Manager API is running!"); err != nil {
		logger.Log.Error("Failed to write response: ", err)
	}
}

// Функция для запуска HTTP-сервера
func startServer() {
	port := ":" + config.AppConfig.Server.Port
	logger.Log.Infof("Server is running at http://localhost%s", port)

	mux := http.NewServeMux()
	mux.HandleFunc("/", handleRoot)

	server := &http.Server{
		Addr:    port,
		Handler: mux,
	}

	// Запуск сервера в отдельной горутине
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Log.Fatal("Server error: ", err)
		}
	}()

	// Канал для обработки сигнала завершения работы
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// Ожидание сигнала завершения
	<-stop
	logger.Log.Info("Shutting down server...")

	// Контекст с таймаутом для корректного завершения сервера
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logger.Log.Warn("Error during server shutdown: ", err)
	} else {
		logger.Log.Info("Server shutdown complete")
	}

	// Закрытие подключения к базе данных
	database.CloseDB()
}

func main() {
	// Инициализация логера
	logger.InitLogger()
	logger.Log.Info("Loading configuration...")
	config.LoadConfig()

	// Подключение к базе данных
	logger.Log.Info("Connecting to the database...")
	database.InitDB()

	// Запуск HTTP-сервера
	startServer()
}
