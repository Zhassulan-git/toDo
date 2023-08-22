package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/user/toDo/config"
	"github.com/user/toDo/config/models"
)

func GetTodos(w http.ResponseWriter, r *http.Request) {
	var tasks []models.Todo
	rows, err := config.SetupDb().Query("SELECT * FROM todos;")
	if err != nil {
		log.Fatal(err)
	}
	//
	for rows.Next() { // Next prepares the next result row for reading with the Scan method.
		var task models.Todo
		if err := rows.Scan(&task.ID, &task.Task_name, &task.Task, &task.Status); err != nil {
			log.Fatal(err)
		}
		tasks = append(tasks, task)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	} //

	//json.NewEncoder(w).Encode(tasks) //приводит к необходимому формату json
	res, _ := json.Marshal(tasks)

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	var task models.Todo

	rows, err := config.SetupDb().Query("SELECT * FROM todos WHERE task_id = $1;", id)
	if err != nil {
		log.Fatal(err)
	}
	rows.Next()
	if err := rows.Scan(&task.ID, &task.Task_name, &task.Task, &task.Status); err != nil { //&task.`field` is the var of our struct created in models package
		log.Fatal(err)
	}
	res, _ := json.Marshal(task)

	w.Header().Set("Content-type", "application-json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
func AddTodos(w http.ResponseWriter, r *http.Request) {
	db := config.SetupDb()
	var task models.Todo
	err := json.NewDecoder(r.Body).Decode(&task)

	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec("INSERT INTO todos (task_id, task_name, task, status) VALUES ($1, $2, $3, $4);", task.ID, task.Task, task.Task_name, task.Status)
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(task)
}

func DeleteById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	db := config.SetupDb()

	id := vars["id"]

	_, err := db.Query("SELECT * FROM todos WHERE task_id = $1;", id)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		log.Fatal(err, "77")
		return
	} else {
		_, err := db.Exec("DELETE FROM todos WHERE task_id = $1;", id)
		if err != nil {
			log.Fatal(err)
		}
	}
	json.NewEncoder(w).Encode("Task deleted")
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	db := config.SetupDb()

	var task models.Todo
	err := json.NewDecoder(r.Body).Decode(&task)

	_, err = db.Query("SELECT * FROM todos WHERE task_id = $1;", id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		log.Fatal(err)
		return
	} else {
		_, err := db.Exec(`UPDATE todos 
			SET task_id = $1, task_name = $2, task = $3, status = $4 
			WHERE task_id = $1;`, task.ID, task.Task, task.Task_name, task.Status)
		if err != nil {
			log.Fatal(err)
		}
	}
	json.NewEncoder(w).Encode("Task Updated")

}
