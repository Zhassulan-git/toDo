package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/user/toDo/config/models"
	"golang.org/x/crypto/bcrypt"
)

var SecretKey = []byte(os.Getenv("SECRET_KEY"))

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		c, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				// If the cookie is not set, return an unauthorized status
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w, "Not Authorized")
				return
			}
			// For any other type of error, return a bad request status
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		tkn, err := models.CheckAuth(c)
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if !tkn.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		next(w, r)
	}
}

func LogInUser(w http.ResponseWriter, r *http.Request) {

	var err error

	// get email and pass
	var person, user models.UserLogIn

	json.NewDecoder(r.Body).Decode(&user)
	email := user.Email
	password := user.Password

	//get datas from database
	row := models.GetFromDB(email, password)
	row.Scan(&person.Password)

	//compare sent request and saved pass hash
	error := bcrypt.CompareHashAndPassword([]byte(person.Password), []byte(password))
	if error != nil {
		log.Fatal(err, "Invalid password")
	} else {
		fmt.Println("passwords equal")
	}

	// Create the JWT claims, which includes the username and expiry time
	expirationTime := time.Now().Add(5 * time.Minute)

	claims := models.AddCliams(email, expirationTime)
	//generate a jwt token
	token := models.GenerateJWTToken(*claims)

	//add signatures for header of request for authentication later
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   token,
		Expires: expirationTime,
	})
}
