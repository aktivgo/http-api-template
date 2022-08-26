package app

import (
	"database/sql"
	"wire-di/internal/app/server"
	"wire-di/internal/user"

	_ "github.com/lib/pq"
)

func Run() {
	db, err := sql.Open("postgres", "")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	userHandler := user.Wire(db)

	server.CreateUserHttpServer("8080", userHandler)
}
