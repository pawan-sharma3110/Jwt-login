package model

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type User struct {
	UserId    int       `json:"user_id"`
	FirstName string    `json:"first_name"`
	MobileNo  string    `json:"mobile_no"`
	EmailId   string    `json:"email_id"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

// JWT Login Models

type Claims struct {
	EmailID string `json:"email_id"`
	jwt.StandardClaims
}