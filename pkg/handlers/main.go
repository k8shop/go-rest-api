package handlers

import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/k8shop/go-rest-api/pkg/informer"
)

//Register all handlers for this API
func Register(router *mux.Router, db *gorm.DB, informer *informer.Informer) error {
	for _, handler := range GetAllHandlers() {
		handler.Register(db, informer, router.PathPrefix("/"+handler.Slug()).Subrouter())
	}
	return nil
}
