package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/k8shop/go-rest-api/pkg/models"
)

//Products handler
type Products struct {
	db *sql.DB
}

//NewProductsHandler handles products handles
func NewProductsHandler() *Products {
	return &Products{}
}

//Slug slog
func (p *Products) Slug() string {
	return "products"
}

//Register is register
func (p *Products) Register(db *sql.DB, router *mux.Router) {
	p.db = db
	router.HandleFunc("/", p.handleGet).Methods("GET")
	router.HandleFunc("/", p.handlePost).Methods("POST")
	router.HandleFunc("/{id:[0-9]+}", p.handleDelete).Methods("DELETE")
}

func (p *Products) handleGet(res http.ResponseWriter, req *http.Request) {
	results, err := p.db.Query("SELECT * FROM products")
	if err != nil {
		panic(err)
	}

	products := []models.Product{}
	for results.Next() {
		product := models.Product{}
		err = results.Scan(&product.ID, &product.Name)
		if err != nil {
			panic(err)
		}
		products = append(products, product)
	}

	resBytes, err := json.Marshal(products)
	if err != nil {
		res.Write([]byte(err.Error()))
		return
	}
	res.Write(resBytes)
}

func (p *Products) handlePost(res http.ResponseWriter, req *http.Request) {
	product := models.Product{
		Name: req.FormValue("name"),
	}

	_, err := p.db.Exec("INSERT INTO products (name) VALUES (?)", product.Name)
	if err != nil {
		res.Write([]byte(err.Error()))
		return
	}
	resBytes, err := json.Marshal(product)
	if err != nil {
		res.Write([]byte(err.Error()))
		return
	}
	res.Write(resBytes)
}

func (p *Products) handleDelete(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	_, err := p.db.Exec("DELETE FROM products WHERE id = ?", params["id"])
	if err != nil {
		res.Write([]byte(err.Error()))
		return
	}
	res.Write([]byte("{\"result\": \"success\"}"))
}
