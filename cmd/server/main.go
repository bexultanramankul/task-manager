package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"task-manager/internal/config"
	"task-manager/internal/delivery"
	"task-manager/internal/repository"
	"task-manager/internal/routes"
	"task-manager/internal/storage"
	"task-manager/internal/usecase"
	"task-manager/pkg/logger"
	"time"
)

// Обработчик корневого маршрута
func handleRoot(w http.ResponseWriter, r *http.Request) {
	if _, err := fmt.Fprintln(w, "Task Manager API is running!"); err != nil {
		logger.Log.Error("Failed to write response: ", err)
	}
}

func main() {
	// Инициализация логера
	logger.InitLogger()
	logger.Log.Info("Loading configuration...")
	config.LoadConfig()

	// Подключение к базе данных
	logger.Log.Info("Connecting to the storage...")
	storage.InitDB()

	taskRepo := repository.NewTaskRepository(storage.DB)
	boardRepo := repository.NewBoardRepository(storage.DB)
	userRepo := repository.NewUserRepository(storage.DB)

	taskUsecase := usecase.NewTaskUsecase(taskRepo)
	boardUsecase := usecase.NewBoardUsecase(boardRepo)
	userUsecase := usecase.NewUserUsecase(userRepo)

	taskHandler := delivery.NewTaskHandler(taskUsecase)
	boardHandler := delivery.NewBoardHandler(boardUsecase)
	userHandler := delivery.NewUserHandler(userUsecase)

	// Создание маршрутизатора
	r := router.NewRouter(taskHandler, boardHandler, userHandler)

	// Добавляем корневой обработчик
	r.HandleFunc("/", handleRoot).Methods("GET")

	// Запуск сервера
	port := ":" + config.AppConfig.Server.Port
	server := &http.Server{
		Addr:    port,
		Handler: r, // Используем маршрутизатор gorilla/mux
	}

	logger.Log.Infof("Server is running at http://localhost%s", port)

	// Запуск сервера в отдельной горутине
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Log.Fatal("Server error: ", err)
		}
	}()

	// Ожидание сигнала завершения
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
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

	// Закрытие базы данных
	storage.CloseDB()
}
