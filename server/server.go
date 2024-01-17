package server

import (
	"net/http"
	"screening/handlers"
)

type ServerHttp struct {
	App http.Handler
}

func ServeHTTP(handlers *handlers.Handler) *ServerHttp {
	mux := http.NewServeMux()
	mux.HandleFunc("/createUser", handlers.CreateUser)
	mux.HandleFunc("/updateUser", handlers.UpdateUser)

	return &ServerHttp{App: mux}
}

func (s *ServerHttp) Start() {
	http.ListenAndServe(":8000", s.App)
}
