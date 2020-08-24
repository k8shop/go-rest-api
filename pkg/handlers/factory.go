package handlers

import (
	"database/sql"

	"github.com/gorilla/mux"
)

//NewHandler by type
func NewHandler(handler string) Interface {
	switch handler {
	case "sample":
		return NewSampleHandler()
	default:
		return &NoOpHandler{}
	}
}

//GetAllHandlers available
func GetAllHandlers() []Interface {
	return []Interface{
		NewSampleHandler(),
		NewBikesHandler(),
	}
}

//NoOpHandler non operational handler
type NoOpHandler struct {
}

//Register nothing
func (n *NoOpHandler) Register(db *sql.DB, router *mux.Router) {
}

//Slug for NoOpHandler
func (n *NoOpHandler) Slug() string {
	return ""
}
