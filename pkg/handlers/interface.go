package handlers

import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/k8shop/go-rest-api/pkg/informer"
)

//Interface for all handlers
type Interface interface {
	Register(db *gorm.DB, informer *informer.Informer, router *mux.Router)
	Slug() string
}
