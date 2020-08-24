package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/k8shop/go-rest-api/pkg/models"
)

//Bikes handler
type Bikes struct {
	db *sql.DB
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
func (b *Bikes) Register(db *sql.DB, router *mux.Router) {
	b.db = db
	router.HandleFunc("/", b.handleGet).Methods("GET")
	router.HandleFunc("/", b.handlePost).Methods("POST")
	router.HandleFunc("/{id:[0-9]+}", b.handleDelete).Methods("DELETE")
}

func (b *Bikes) handleGet(res http.ResponseWriter, req *http.Request) {
	results, err := b.db.Query("SELECT * FROM bikes")
	if err != nil {
		panic(err)
	}

	samples := []models.Sample{}
	for results.Next() {
		sample := models.Sample{}
		err = results.Scan(&sample.ID, &sample.Name)
		if err != nil {
			panic(err)
		}
		samples = append(samples, sample)
	}

	resBytes, err := json.Marshal(samples)
	if err != nil {
		res.Write([]byte(err.Error()))
		return
	}
	res.Write(resBytes)
}

func (b *Bikes) handlePost(res http.ResponseWriter, req *http.Request) {
	sample := models.Sample{
		Name: req.FormValue("name"),
	}

	_, err := b.db.Exec("INSERT INTO bikes (name) VALUES (?)", sample.Name)
	if err != nil {
		res.Write([]byte(err.Error()))
		return
	}
	resBytes, err := json.Marshal(sample)
	if err != nil {
		res.Write([]byte(err.Error()))
		return
	}
	res.Write(resBytes)
}

func (b *Bikes) handleDelete(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	_, err := b.db.Exec("DELETE FROM bikes WHERE id = ?", params["id"])
	if err != nil {
		res.Write([]byte(err.Error()))
		return
	}
	res.Write([]byte("{\"result\": \"success\"}"))
}
