package models

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/user/toDo/config"
)

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

var SecretKey = []byte(os.Getenv("SECRET_KEY"))

type Claims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

//function to generate new tokens
func GenerateJWTToken(claims Claims) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(SecretKey)

	if err != nil {
		fmt.Errorf("Something went wrong: %s", err.Error())
	}
	return tokenString
}

//get valeus from cokie and parse jwt token
func CheckAuth(c *http.Cookie) (*jwt.Token, error) {

	// Get the JWT string from the cookie
	tknStr := c.Value

	// Initialize a new instance of `Claims`
	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})

	return tkn, err
}

//USER AUTHENTICATION
func GetFromDB(email string, pass string) *sql.Row {

	db := config.SetupDb()

	query := "SELECT password FROM users WHERE email = $1;"
	row := db.QueryRow(query, email)

	return row
}

func AddCliams(email string, expirationTime time.Time) *Claims {

	claims := &Claims{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	return claims
}
