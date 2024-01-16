package main

import (
	"log"
	"net/http"
)

func init() {
	// create mysql connection
	db, err := createConnection()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	DB = db
}

func main() {
	// setup json api
	setupJsonApi()

	// start server
	http.ListenAndServe(":8000", nil)
}
