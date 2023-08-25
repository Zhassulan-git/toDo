package models

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func GeneratePasswordHash(pass string) []byte {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), 14)
	if err != nil {
		log.Fatal(err, "Failed to hash password")
	}
	return bytes
}
