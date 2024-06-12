package utils

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"jwt/model"
	"log"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func ParseJson(r *http.Request, payload any) error {
	if r.Body == nil {
		err := fmt.Errorf("enter valid paylod in request body ")
		return err
	}
	err := json.NewDecoder(r.Body).Decode(payload)
	if err != nil {
		return fmt.Errorf("internal server error")

	}
	defer r.Body.Close()
	return nil
}
func WriteJson(w http.ResponseWriter, status int, res any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}

	return nil
}
func WriteError(w http.ResponseWriter, status int, err error) {
	WriteJson(w, status, map[string]string{"error": err.Error()})
}
func createUSerTable(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS users(
		user_id SERIAL PRIMARY KEY,
		first_name TEXT NOT NULL,
		mobile_no TEXT NOT NULL,
		email_id TEXT NOT NULL,
		password TEXT NOT NULL,
		created_at  TIMESTAMP
)`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatalf("Falied in to create user table :%v", err)
		return err
	}

	return nil
}
func isUserExists(db *sql.DB, user model.User) (bool, error) {
	var email string
	query := `SELECT email_id FROM users WHERE email_id=$1`
	err := db.QueryRow(query, user.EmailId).Scan(&email)

	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	return true, nil
}
func InsertUser(db *sql.DB, user model.User) (int, error) {
	err := createUSerTable(db)
	if err != nil {
		return 0, fmt.Errorf("failed to create user table: %v", err)
	}
	exists, err := isUserExists(db, user)
	if err != nil {
		return 0, err
	}
	if exists {
		return 0, fmt.Errorf("user already exists")
	}
	hashPass, err := gernateHashPass(user)
	if err != nil {
		return 0, err
	}
	var userId int
	query := `INSERT INTO users(first_name, mobile_no, email_id, password, created_at) VALUES($1, $2, $3, $4, $5) RETURNING user_id`
	err = db.QueryRow(query, user.FirstName, user.MobileNo, user.EmailId, hashPass, time.Now()).Scan(&userId)
	if err != nil {
		return 0, fmt.Errorf("failed to insert user into table: %v", err)
	}
	return userId, nil
}
func gernateHashPass(user model.User) (string, error) {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		err = errors.New("falied to convert password in hash")
		return "", err
	}
	return string(hashPass), nil
}


func UserLogin(db *sql.DB) {

}
