package routes

import (
	"github.com/gorilla/mux"
	"github.com/user/toDo/config/controllers"
)

// func RegisterTasks(r *mux.Router) {
// 	r.Handle("/", controllers.AuthMiddleware(controllers.GetTodos)).Methods("GET")
// 	r.Handle("/{id}", controllers.AuthMiddleware(controllers.GetById)).Methods("GET")
// 	r.Handle("/", controllers.AuthMiddleware(controllers.AddTodos)).Methods("POST")
// 	r.Handle("/{id}", controllers.AuthMiddleware(controllers.DeleteById)).Methods("DELETE")
// 	r.Handle("/{id}", controllers.AuthMiddleware(controllers.UpdateTask)).Methods("PUT")
// }

func RegisterTasks(r *mux.Router) {
	r.HandleFunc("/", controllers.GetTodos).Methods("GET")
	r.HandleFunc("/{id}", controllers.GetById).Methods("GET")
	r.HandleFunc("/", controllers.AddTodos).Methods("POST")
	r.HandleFunc("/{id}", controllers.DeleteById).Methods("DELETE")
	r.HandleFunc("/{id}", controllers.UpdateTask).Methods("PUT")
}
