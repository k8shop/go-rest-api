package handlers

import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/k8shop/go-rest-api/pkg/informer"
)

//NewHandler by type
func NewHandler(handler string) Interface {
	switch handler {
	case "products":
		return NewProductsHandler()
	case "registration":
		return NewRegistrationHandler()
	default:
		return &NoOpHandler{}
	}
}

//GetAllHandlers available
func GetAllHandlers() []Interface {
	return []Interface{
		NewProductsHandler(),
		NewRegistrationHandler(),
	}
}

//NoOpHandler non operational handler
type NoOpHandler struct {
}

//Register nothing
func (n *NoOpHandler) Register(_ *gorm.DB, _ *informer.Informer, _ *mux.Router) {
}

//Slug for NoOpHandler
func (n *NoOpHandler) Slug() string {
	return ""
}
