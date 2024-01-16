package main

import (
	"log"
	"net/http"
)

// HandleApiError handles api error
func HandleApiError(w http.ResponseWriter, err error, msg string, code int) {
	log.Printf("Error : %v", err)
	w.WriteHeader(code)
	w.Write([]byte(msg + err.Error()))
}
