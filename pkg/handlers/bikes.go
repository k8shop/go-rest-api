package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/k8shop/go-rest-api/pkg/models"
)

//Bikes handler
type Bikes struct {
}

//NewBikesHandler handles bikes handles
func NewBikesHandler() *Bikes {
	return &Bikes{}
}

//Slug slog
func (b *Bikes) Slug() string {
	return "bikes"
}

//Register is register
func (b *Bikes) Register(router *mux.Router) {
	router.HandleFunc("/", b.handleGet)
}

func (b *Bikes) handleGet(res http.ResponseWriter, req *http.Request) {
	samples := []*models.Sample{
		&models.Sample{
			ID:   1,
			Name: "a bike, wooo",
		},
	}

	resBytes, err := json.Marshal(samples)
	if err != nil {
		res.Write([]byte(err.Error()))
		return
	}
	res.Write(resBytes)
}
