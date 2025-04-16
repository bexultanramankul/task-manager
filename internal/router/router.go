package router

import (
	"github.com/gin-gonic/gin"
	"task-manager/internal/handler"
	"task-manager/internal/middleware"
)

func SetupRouter(
	userHandler *handler.UserHandler,
	taskHandler *handler.TaskHandler,
	boardHandler *handler.BoardHandler,
	jwtSecret string,
) *gin.Engine {
	r := gin.Default()

	// Инициализация middleware
	authMiddleware := middleware.AuthMiddleware(jwtSecret)
	accessMiddleware := middleware.NewAccessMiddleware(jwtSecret)

	// Группа публичных маршрутов (не требующих аутентификации)
	public := r.Group("/api")
	{
		public.POST("/register", userHandler.RegisterUser)
		public.POST("/login", userHandler.Login)
	}

	// Группа защищенных маршрутов (требующих аутентификации)
	auth := r.Group("/api")
	auth.Use(authMiddleware)
	{
		// Пользовательские маршруты
		//auth.GET("/users/me", userHandler.GetCurrentUser)
		//auth.PUT("/users/me", userHandler.UpdateCurrentUser)
		//auth.DELETE("/users/me", userHandler.DeleteCurrentUser)

		// Маршруты досок
		auth.GET("/boards", boardHandler.GetAllBoards)
		auth.POST("/boards", boardHandler.CreateBoard)

		// Маршруты для конкретной доски
		board := auth.Group("/boards/:id")
		board.Use(accessMiddleware.BoardOwner())
		{
			board.GET("", boardHandler.GetBoard)
			board.PUT("", boardHandler.UpdateBoard)
			board.DELETE("", boardHandler.DeleteBoard)
			board.POST("/block", boardHandler.BlockBoard)

			// Маршруты задач для конкретной доски
			//board.GET("/tasks", taskHandler.GetTasksByBoard)
			board.POST("/tasks", taskHandler.CreateTask)
		}

		// Маршруты для конкретной задачи
		task := auth.Group("/tasks/:id")
		task.Use(accessMiddleware.TaskOwnerOrBoardOwner())
		{
			task.GET("", taskHandler.GetTask)
			task.PUT("", taskHandler.UpdateTask)
			task.DELETE("", taskHandler.DeleteTask)
			task.POST("/block", taskHandler.BlockTask)
		}

		// Админские маршруты
		admin := auth.Group("/admin")
		admin.Use(accessMiddleware.AdminOnly())
		{
			admin.GET("/users", userHandler.GetAllUsers)
			admin.GET("/users/:id", userHandler.GetUser)
			admin.PUT("/users/:id", userHandler.UpdateUser)
			admin.DELETE("/users/:id", userHandler.DeleteUser)
		}
	}

	return r
}
