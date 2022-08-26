package server

import (
	"wire-di/internal/domain"

	"github.com/gorilla/mux"
)

func buildRouter(
	userHandler domain.UserHandler,
) *mux.Router {
	router := mux.NewRouter()

	router.Handle("/users", userHandler.FetchByUsername())

	return router
}
