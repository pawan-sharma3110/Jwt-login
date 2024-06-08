package db

import (
	"database/sql"
	"log"
)

func DbIn() (db *sql.DB, err error) {
	conStr := `host=localhost port=5432 username=postgres dbname=Jwt password=Pawan@123 sslmode=disable`
	db, err = sql.Open("postgres", conStr)
	if err != nil {
		log.Fatalf("Database Error : %v", err)
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("Database connection Error : %v", err)
		return nil, err
	}
	return db, nil
}

