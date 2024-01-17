package main

import (
	"screening/di"
)

// func init() {
// 	// create mysql connection
// 	db, err := createConnection()
// 	if err != nil {
// 		log.Fatalf("Error connecting to database: %v", err)
// 	}
// 	db.SetMaxOpenConns(20)
// 	db.SetMaxIdleConns(10)
// 	DB = db
// }

func main() {
	server := di.InitDi()
	server.Start()
}
