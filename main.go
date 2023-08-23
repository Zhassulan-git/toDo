package main

import (
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq" //for unknown driver "postgres"
	"github.com/user/toDo/config/routes"
)

func main() {
	r := mux.NewRouter()

	routes.RegisterTasks(r)
	routes.UsersRegister(r)

	http.Handle("/", r)
	http.ListenAndServe(":8000", nil)
}
