package main

import (
	"jwt/db"
	"jwt/router"
	"log"
	"net/http"
)

func main() {
	db, err := db.DbIn()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := router.MyRoutes()
	http.ListenAndServe(":8080", r)
}
