package routes

import (
	"github.com/gorilla/mux"
	controllers "github.com/user/toDo/config/controllers/tasks"
	"github.com/user/toDo/config/service"
)

func RegisterTasks(r *mux.Router) {
	r.Handle("/", service.CheckAuth(controllers.GetTodos)).Methods("GET")
	r.Handle("/{id}", service.CheckAuth(controllers.GetById)).Methods("GET")
	r.Handle("/", service.CheckAuth(controllers.AddTodos)).Methods("POST")
	r.Handle("/{id}", service.CheckAuth(controllers.DeleteById)).Methods("DELETE")
	r.Handle("/{id}", service.CheckAuth(controllers.UpdateTask)).Methods("PUT")
}
