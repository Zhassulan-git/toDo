package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/user/toDo/config"
	"github.com/user/toDo/config/models"
)

func SignUpUser(w http.ResponseWriter, r *http.Request) {

	var person models.User

	var err error
	err = json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		log.Fatal(err)
	}

	db := config.SetupDb()

	var datafield models.UserLogIn

	err = json.Unmarshal(person.UserLog, &datafield)

	if err != nil {
		log.Fatal(err, "unmarshal fatal")
	}

	bytes := models.GeneratePasswordHash(datafield.Password)

	_, err = db.Exec(`INSERT INTO 
			users (personid, firstname, lastname, email, password) 
			VALUES ($1, $2, $3, $4, $5);`,
		person.ID, person.First_Name, person.Last_Name, datafield.Email, bytes)

	if err != nil {
		log.Fatal(err, "sql exec failed")
	}

	json.NewEncoder(w).Encode(person)
}
