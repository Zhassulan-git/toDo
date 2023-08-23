package routes

import (
	"github.com/gorilla/mux"
	"github.com/user/toDo/config/controllers"
)

func UsersRegister(r *mux.Router) {
	r.HandleFunc("/add", controllers.SignUpUser).Methods("POST")

	r.HandleFunc("/login", controllers.LogInUser).Methods("POST")
}
