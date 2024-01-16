package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// createConnection creates a connection to mysql database
func createConnection() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:my-secret-pw@tcp(localhost:3306)/test")
	if err != nil {
		return nil, err
	}

	qury := `CREATE TABLE IF NOT EXISTS users (
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(50) NOT NULL,
		email VARCHAR(50) NOT NULL UNIQUE
		);`

	_, err = db.Exec(qury)
	if err != nil {
		return nil, err
	}
	return db, nil
}
