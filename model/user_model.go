package model

import "time"

type User struct {
	UserId    int       `json:"user_id"`
	FirstName string    `json:"first_name"`
	MobileNo  string    `json:"mobile_no"`
	EmailId   string    `json:"email_id"`
	CreatedAt time.Time `json:"created_at"`
}
