package router

import (
	"github.com/gorilla/mux"
	"github.com/prajapatiomkar/todo-app-in-golang/middleware"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/task", middleware.GetAllTask).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/task", middleware.CreateTask).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/task/{id}", middleware.TaskComplete).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/undo-task/{id}", middleware.UndoTask).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/delete-task/{id}", middleware.DeleteTask).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/api/delete-all-task/", middleware.DeleteAllTask).Methods("DELETE", "OPTIONS")

	return router
}
