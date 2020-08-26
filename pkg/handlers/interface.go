package handlers

import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

//Interface for all handlers
type Interface interface {
	Register(db *gorm.DB, router *mux.Router)
	Slug() string
}
