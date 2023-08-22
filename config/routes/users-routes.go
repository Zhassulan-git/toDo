package routes

import (
	"github.com/gorilla/mux"
	"github.com/user/toDo/config/service"
)

func UsersRegister(r *mux.Router) {
	r.HandleFunc("/add", service.SignUpUser).Methods("POST")

	r.HandleFunc("/login", service.LogInUser).Methods("POST")
}
