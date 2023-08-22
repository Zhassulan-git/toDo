package main

import (
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq" //for unknown driver "postgres"
	"github.com/user/toDo/config/routes"
)

func main() {
	r := mux.NewRouter()

	//task := r.PathPrefix("/todos").Subrouter()
	routes.RegisterTasks(r)

	//user := r.PathPrefix("/users").Subrouter()

	//user.Use(controllers.AuthMiddleware)
	routes.UsersRegister(r)

	http.Handle("/", r)
	http.ListenAndServe(":8000", nil)
}
