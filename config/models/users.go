package models

import "encoding/json"

type User struct {
	ID         int             `json:"id"`
	First_Name string          `json:"first_name"`
	Last_Name  string          `json:"last_name"`
	UserLog    json.RawMessage `json:"userlog"`
}

type UserLogIn struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
