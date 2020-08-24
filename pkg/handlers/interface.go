package handlers

import (
	"database/sql"

	"github.com/gorilla/mux"
)

//Interface for all handlers
type Interface interface {
	Register(db *sql.DB, router *mux.Router)
	Slug() string
}
