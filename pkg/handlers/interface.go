package handlers

import "github.com/gorilla/mux"

//Interface for all handlers
type Interface interface {
	Register(router *mux.Router)
	Slug() string
}
