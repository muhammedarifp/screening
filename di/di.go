package di

import (
	"screening/db"
	"screening/handlers"
	"screening/repo"
	"screening/server"
	"screening/usecases"
)

func InitDi() *server.ServerHttp {
	db, dbErr := db.CreateConnection()
	if dbErr != nil {
		panic(dbErr)
	}

	repo := repo.New(db)
	useCase := usecases.New(repo)
	handlers := handlers.New(useCase)
	server := server.ServeHTTP(handlers)
	return server
}
