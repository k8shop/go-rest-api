package handlers

import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

//NewHandler by type
func NewHandler(handler string) Interface {
	switch handler {
	case "products":
		return NewProductsHandler()
	default:
		return &NoOpHandler{}
	}
}

//GetAllHandlers available
func GetAllHandlers() []Interface {
	return []Interface{
		NewProductsHandler(),
	}
}

//NoOpHandler non operational handler
type NoOpHandler struct {
}

//Register nothing
func (n *NoOpHandler) Register(db *gorm.DB, router *mux.Router) {
}

//Slug for NoOpHandler
func (n *NoOpHandler) Slug() string {
	return ""
}
