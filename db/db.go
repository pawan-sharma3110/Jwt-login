package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func DbIn() (db *sql.DB, err error) {
	conStr := `host=localhost port=5432 user= postgres password=Pawan@2003 dbname=Jwt sslmode=disable`
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
