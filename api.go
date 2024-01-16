package main

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
)

// every handlers call connectdb function is not good approach
// so we create a global variable
var (
	DB *sql.DB
)

func setupJsonApi() {
	http.HandleFunc("/createUser", func(w http.ResponseWriter, r *http.Request) {
		// create mysql connection

		name := r.FormValue("name")
		email := r.FormValue("email")
		query := "INSERT INTO users (name, email) VALUES (?,?)"

		// check params is not empty
		if name == "" || email == "" {
			HandleApiError(w, errors.New("name or email is empty"), "email and password is required", http.StatusBadGateway)
			return
		}

		// check email is already exist
		var count int
		if err := DB.QueryRow(`SELECT COUNT(*) FROM users WHERE email = ?`, email).Scan(&count); err != nil {
			HandleApiError(w, err, "Error creating user. Please try again! ", http.StatusBadGateway)
			return
		}

		if count > 0 {
			HandleApiError(w, nil, "Email already exist", http.StatusBadGateway)
			return
		}

		// insert user
		result, err := DB.Exec(query, name, email)

		// if there is an error inserting, handle it
		if err != nil {
			HandleApiError(w, err, "Error creating user. Please try again! ", http.StatusBadGateway)
			return
		}

		// get the id of the inserted row
		id, iderr := result.LastInsertId()
		if iderr != nil {
			HandleApiError(w, iderr, "Error creating user. Please try again! ", http.StatusBadGateway)
			return
		}
		fmt.Println("Last inserted ID: ", id)
		w.Write([]byte("Created user successfully!"))
	})

	http.HandleFunc("/updateUser", func(w http.ResponseWriter, r *http.Request) {
		// parse form data
		if err := r.ParseForm(); err != nil {
			w.WriteHeader(400)
			w.Write([]byte("Error parsing form. Please try again! " + err.Error()))
			return
		}

		// get form data
		name := r.FormValue("name")
		email := r.FormValue("email")
		id := r.FormValue("id")

		// check params is not empty
		if name == "" || email == "" || id == "" {
			w.WriteHeader(400)
			w.Write([]byte("email,userid and password is required"))
			return
		}

		// update user
		query := "UPDATE users SET name= ?, email = ? WHERE id = ?"
		result, err := DB.Exec(query, name, email, id)
		if err != nil {
			w.WriteHeader(400)
			w.Write([]byte("Error updating user. Please try again! " + err.Error()))
			return
		}

		// get the number of rows affected by the update
		fmt.Println("Updated ID : ", result)
		w.Write([]byte("User updated successfully!"))
	})
}
