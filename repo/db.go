package repo

import (
	"database/sql"
	"fmt"
)

type DataBase struct {
	DB *sql.DB
}

func New(db *sql.DB) *DataBase {
	return &DataBase{
		DB: db,
	}
}

func (d *DataBase) CreateNewUser(name, email string) (int64, error) {
	fmt.Println("name: ", name)
	fmt.Println("email: ", email)
	query := "INSERT INTO users (name, email) VALUES (?, ?)"
	result, err := d.DB.Exec(query, name, email)
	if err != nil {
		return 0, fmt.Errorf("error while creating new user: %w", err)
	}

	// get the last inserted id
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("error while getting last inserted id: %w", err)
	}

	return id, nil
}

func (d *DataBase) UpdateUser(id int64, name, email string) (int64, error) {
	query := "UPDATE users SET name = ?, email = ? WHERE id = ?"
	result, err := d.DB.Exec(query, name, email, id)
	if err != nil {
		return 0, fmt.Errorf("error while updating user: %w", err)
	}

	// get the number of rows affected by the insert
	rows, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("error while getting rows affected: %w", err)
	}

	return rows, nil
}

func (d *DataBase) Close() error {
	return d.DB.Close()
}
