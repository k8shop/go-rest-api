package handlers

import (
	"github.com/gorilla/mux"
)

//Register all handlers for this API
func Register(router *mux.Router) error {
	for _, handler := range GetAllHandlers() {
		handler.Register(router.PathPrefix("/" + handler.Slug()).Subrouter())
	}
	return nil
}
