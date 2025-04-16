package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"task-manager/internal/config"
	"task-manager/internal/handler"
	"task-manager/internal/model"
	"task-manager/internal/repository"
	"task-manager/internal/router"
	"task-manager/internal/storage"
	"task-manager/internal/usecase"
	"task-manager/pkg/logger"
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

	if err := config.LoadConfig("./configs"); err != nil {
		logger.Log.Fatal("Failed to load config: ", err)
	}

	cfg := config.GetConfig()

	// Подключение к базе данных
	logger.Log.Info("Connecting to the storage...")
	storage.InitDB()
	db := storage.DB

	// Auto migrate models
	logger.Log.Info("Running auto migrations...")
	if err := db.AutoMigrate(
		&model.User{},
		&model.Board{},
		&model.Task{},
	); err != nil {
		logger.Log.Fatal("Failed to auto migrate models: ", err)
	}

	// 5. Инициализация репозиториев
	logger.Log.Info("Initializing repositories...")
	boardRepo := repository.NewBoardRepository(db)
	taskRepo := repository.NewTaskRepository(db)
	userRepo := repository.NewUserRepository(db)

	// 6. Инициализация use cases
	logger.Log.Info("Initializing use cases...")
	boardUsecase := usecase.NewBoardUsecase(boardRepo)
	taskUsecase := usecase.NewTaskUsecase(taskRepo)
	userUsecase := usecase.NewUserUsecase(userRepo)

	// 7. Инициализация обработчиков
	logger.Log.Info("Initializing handlers...")
	boardHandler := handler.NewBoardHandler(boardUsecase)
	taskHandler := handler.NewTaskHandler(taskUsecase)
	userHandler := handler.NewUserHandler(userUsecase, cfg.Auth.JWTSecret)

	// 8. Настройка маршрутизатора
	logger.Log.Info("Setting up router...")
	r := router.SetupRouter(
		userHandler,
		taskHandler,
		boardHandler,
		config.GetConfig().Auth.JWTSecret,
	)

	// 9. Настройка корневого маршрута
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Task Manager API is running!")
	})

	// 10. Запуск сервера
	serverConfig := cfg.Server
	logger.Log.Infof("Starting server on %s:%s", serverConfig.Host, serverConfig.Port)
	if err := r.Run(fmt.Sprintf("%s:%s", serverConfig.Host, serverConfig.Port)); err != nil {
		logger.Log.Fatal("Failed to start server: ", err)
	}
}
