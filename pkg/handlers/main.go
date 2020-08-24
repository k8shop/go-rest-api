package handlers

import (
	"database/sql"

	"github.com/gorilla/mux"
)

//Register all handlers for this API
func Register(router *mux.Router, db *sql.DB) error {
	for _, handler := range GetAllHandlers() {
		handler.Register(db, router.PathPrefix("/"+handler.Slug()).Subrouter())
	}
	return nil
}
