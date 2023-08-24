package controllers

import (
	"os"
)

var SecretKey = []byte(os.Getenv("SECRET_KEY"))
