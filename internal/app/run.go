package app

import (
	"database/sql"
	"net/http"
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
	http.Handle("/user", userHandler.FetchByUsername())
	http.ListenAndServe(":8000", nil)
}
