package router

import (
	"task-manager/internal/delivery"

	"github.com/gorilla/mux"
)

func NewRouter(taskHandler *delivery.TaskHandler, boardHandler *delivery.BoardHandler, userHandler *delivery.UserHandler) *mux.Router {
	r := mux.NewRouter()

	// Task router
	r.HandleFunc("/tasks", taskHandler.GetAllTasks).Methods("GET")
	r.HandleFunc("/tasks/{id:[0-9]+}", taskHandler.GetTaskByID).Methods("GET")
	r.HandleFunc("/tasks", taskHandler.CreateTask).Methods("POST")
	r.HandleFunc("/tasks/{id:[0-9]+}", taskHandler.UpdateTask).Methods("PUT")
	r.HandleFunc("/tasks/{id:[0-9]+}", taskHandler.DeleteTask).Methods("DELETE")

	// Board router
	r.HandleFunc("/boards", boardHandler.GetAllBoards).Methods("GET")
	r.HandleFunc("/boards/{id:[0-9]+}", boardHandler.GetBoardByID).Methods("GET")
	r.HandleFunc("/boards", boardHandler.CreateBoard).Methods("POST")
	r.HandleFunc("/boards/{id:[0-9]+}", boardHandler.UpdateBoard).Methods("PUT")
	r.HandleFunc("/boards/{id:[0-9]+}", boardHandler.DeleteBoard).Methods("DELETE")

	// User router
	r.HandleFunc("/users", userHandler.GetAllUsers).Methods("GET")
	r.HandleFunc("/users/{id:[0-9]+}", userHandler.GetUserByID).Methods("GET")
	r.HandleFunc("/register", userHandler.RegisterUser).Methods("POST")
	r.HandleFunc("/users/{id:[0-9]+}", userHandler.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id:[0-9]+}", userHandler.DeleteUser).Methods("DELETE")

	return r
}
