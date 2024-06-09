package utils

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"jwt/model"
	"log"
	"net/http"
	"time"
)

func ParseJson(r *http.Request, payload any) error {
	if r.Body == nil {
		err := fmt.Errorf("enter valid paylod in request body ")
		return err
	}
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(payload)
}
func WriteJson(w http.ResponseWriter, status int, res any) error {
	w.Header().Set("Content-Type", "application/json") // Correct the content-type typo
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(res)
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
	res, err := db.Exec(query)
	if err != nil {
		log.Fatalf("Falied in to create user table :%v", err)
		return err
	}
	if res != nil {
		println("Table crated")
	}
	return nil
}
func IsertUser(db *sql.DB, w http.ResponseWriter, user model.User) (int, error) {
	err := createUSerTable(db)
	if err != nil {
		http.Error(w, "Falied in to create user table", http.StatusInternalServerError)
		return 0, err
	}
	var userId int
	query := `INSERT INTO users(first_name,mobile_no,email_id,password,created_at)VALUES($1,$2,$3,$4,$5)RETURNING user_id`
	err = db.QueryRow(query, user.FirstName, user.MobileNo, user.EmailId, user.Password, time.Now()).Scan(&userId)
	if err != nil {
		http.Error(w, "Falied to insert user into table", http.StatusInternalServerError)
		return 0, err
	}
	return userId, nil
}
