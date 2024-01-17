package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"screening/helpers"
	"screening/usecases"
	"strconv"
)

// Handler struct
type Handler struct {
	UseCase usecases.UseCaseInterface
}

// New returns new Handler instance
func New(useCase usecases.UseCaseInterface) *Handler {
	return &Handler{
		UseCase: useCase,
	}
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// get form data
		name := r.FormValue("name")
		email := r.FormValue("email")

		fmt.Println("name: ", name)
		fmt.Println("email: ", email)

		// check params is not empty
		if name == "" || email == "" {
			helpers.HandleApiError(w, errors.New("name or email is empty"), "name and email is required", 400)
			return
		}

		// insert user
		id, err := h.UseCase.CreateUser(name, email)
		if err != nil {
			helpers.HandleApiError(w, err, "Error creating user. Please try again! ", 400)
			return
		}

		// return success response
		w.WriteHeader(200)
		w.Write([]byte("User created successfully. ID: " + strconv.FormatInt(id, 10)))
	} else {
		helpers.HandleApiError(w, errors.New(r.Method+" not allowed"), "Only POST method is allowed", 400)
	}
}

func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPut {
		// get form data
		id := r.FormValue("id")
		name := r.FormValue("name")
		email := r.FormValue("email")

		// check params is not empty
		if id == "" || name == "" || email == "" {
			helpers.HandleApiError(w, errors.New("id,name or email is empty"), "id, name and email is required", 400)
			return
		}

		// check id is integer
		idInt, parseErr := strconv.Atoi(id)
		if parseErr != nil {
			helpers.HandleApiError(w, parseErr, "id must be integer", 400)
			return
		}

		// update user
		_, err := h.UseCase.UpdateUser(int64(idInt), name, email)
		if err != nil {
			helpers.HandleApiError(w, err, "Error updating user. Please try again! ", 400)
			return
		}

		// return success response
		w.WriteHeader(200)
		w.Write([]byte("User updated successfully. ID: " + string(id)))
	} else {
		helpers.HandleApiError(w, errors.New(r.Method+" not allowed"), "Only PUT method is allowed", 400)
	}
}
