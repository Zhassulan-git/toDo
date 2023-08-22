package controllers

import (
	"os"
)

var SecretKey = []byte(os.Getenv("SECRET_KEY"))

// func CheckAuth(endpoint func(w http.ResponseWriter, r *http.Request)) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

// 		c, err := r.Cookie("token")
// 		if err != nil {
// 			if err == http.ErrNoCookie {
// 				// If the cookie is not set, return an unauthorized status
// 				w.WriteHeader(http.StatusUnauthorized)
// 				fmt.Fprintf(w, "Not Authorized")
// 				return
// 			}
// 			// For any other type of error, return a bad request status
// 			w.WriteHeader(http.StatusBadRequest)
// 			return
// 		}
// 		// Get the JWT string from the cookie
// 		tknStr := c.Value

// 		// Initialize a new instance of `Claims`
// 		claims := &Claims{}

// 		tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
// 			return SecretKey, nil
// 		})
// 		if err != nil {
// 			if err == jwt.ErrSignatureInvalid {
// 				w.WriteHeader(http.StatusUnauthorized)
// 				return
// 			}
// 			w.WriteHeader(http.StatusBadRequest)
// 			return
// 		}
// 		if !tkn.Valid {
// 			w.WriteHeader(http.StatusUnauthorized)
// 			return
// 		} else {
// 			endpoint(w, r)
// 		}
// 	})
// }
