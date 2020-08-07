package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/k8shop/go-rest-api/pkg/models"
)

//Sample handler
type Sample struct {
}

//NewSampleHandler creates a new sample handler
func NewSampleHandler() *Sample {
	return &Sample{}
}

//Slug for sample handler
func (s *Sample) Slug() string {
	return "sample"
}

//Register routes for this handler
func (s *Sample) Register(router *mux.Router) {
	router.HandleFunc("/", s.handlePost).Methods("POST")
	router.HandleFunc("/", s.handleGet).Methods("GET")
	router.HandleFunc("", s.handleGet).Methods("GET")
}

func (s *Sample) handlePost(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("test response POST"))
}

func (s *Sample) handleGet(res http.ResponseWriter, req *http.Request) {
	samples := []*models.Sample{
		&models.Sample{
			ID:   1,
			Name: "number one",
		},
	}

	resBytes, err := json.Marshal(samples)
	if err != nil {
		res.Write([]byte(err.Error()))
		return
	}
	res.Write(resBytes)
}
