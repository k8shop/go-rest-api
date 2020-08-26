package handlers

import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

//Register all handlers for this API
func Register(router *mux.Router, db *gorm.DB) error {
	for _, handler := range GetAllHandlers() {
		handler.Register(db, router.PathPrefix("/"+handler.Slug()).Subrouter())
	}
	return nil
}
