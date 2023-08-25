package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/user/toDo/config/routes"
)

func main() {
	r := mux.NewRouter()

	routes.UsersRegister(r)
	routes.RegisterTasks(r)

	http.Handle("/", r)
	http.ListenAndServe(":8000", nil)
}
