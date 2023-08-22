package service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/user/toDo/config"
	"github.com/user/toDo/config/models"
	"golang.org/x/crypto/bcrypt"
)

var SecretKey = []byte("83q6$LsD3HAwD0R#rd34")

type Claims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func GenerateJWTToken(claims Claims) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(SecretKey)

	if err != nil {
		fmt.Errorf("Something went wrong: %s", err.Error())
	}
	return tokenString
}

func CheckAuth(endpoint func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

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
		// Get the JWT string from the cookie
		tknStr := c.Value

		// Initialize a new instance of `Claims`
		claims := &Claims{}

		tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
			return SecretKey, nil
		})
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
		} else {
			endpoint(w, r)
		}
	})
}

//USER AUTHENTICATION
func LogInUser(w http.ResponseWriter, r *http.Request) {

	var err error

	// get email and pass

	var person models.UserLogIn

	var user models.UserLogIn

	json.NewDecoder(r.Body).Decode(&user)
	email := user.Email
	password := user.Password

	fmt.Println(email)
	fmt.Println(password)
	db := config.SetupDb()

	query := "SELECT password FROM users WHERE email = $1;"
	err = db.QueryRow(query, email).Scan(&person.Password)

	if err != nil {
		log.Fatal(err, "executing sql query")
	}

	//compare sent request and saved pass hash

	error := bcrypt.CompareHashAndPassword([]byte(person.Password), []byte(password))
	if error != nil {
		log.Fatal(err, "Invalid password")
	} else {
		fmt.Println("passwords equal")
	}

	expirationTime := time.Now().Add(5 * time.Minute)

	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	//generate a jwt token
	token := GenerateJWTToken(*claims)

	//send it back

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   token,
		Expires: expirationTime,
	})
}
