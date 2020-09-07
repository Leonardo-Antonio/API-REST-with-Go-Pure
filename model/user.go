package model

import (
	"github.com/dgrijalva/jwt-go"
)

// User .
type User struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Claim -> datos del payload del token
type Claim struct {
	Email string `json:"email"`
	jwt.StandardClaims
}
