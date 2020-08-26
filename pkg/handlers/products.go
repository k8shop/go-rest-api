package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/k8shop/go-rest-api/pkg/models"
)

//Products handler
type Products struct {
	db *gorm.DB
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
func (p *Products) Register(db *gorm.DB, router *mux.Router) {
	p.db = db
	router.HandleFunc("/", p.handleGet).Methods("GET")
	router.HandleFunc("/{id:[0-9]+}", p.handleGetProduct).Methods("GET")
	router.HandleFunc("/{id:[0-9]+}", p.handlePutProduct).Methods("PUT")
	router.HandleFunc("/", p.handlePost).Methods("POST")
	router.HandleFunc("/{id:[0-9]+}", p.handleDelete).Methods("DELETE")
}

func (p *Products) handleGet(res http.ResponseWriter, req *http.Request) {
	products := []*models.Product{}
	errs := p.db.Debug().Find(&products).GetErrors()
	for _, err := range errs {
		log.Println(err)
	}

	resBytes, err := json.Marshal(products)
	if err != nil {
		res.Write([]byte(err.Error()))
		return
	}
	res.Write(resBytes)
}

func (p *Products) handleGetProduct(res http.ResponseWriter, req *http.Request) {
	product := models.Product{}
	params := mux.Vars(req)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		panic(err)
	}

	errs := p.db.Debug().Find(&product, id).GetErrors()
	for _, err := range errs {
		log.Println(err)
	}

	resBytes, err := json.Marshal(product)
	if err != nil {
		res.Write([]byte(err.Error()))
		return
	}
	res.Write(resBytes)
}

func (p *Products) handlePost(res http.ResponseWriter, req *http.Request) {
	price, err := strconv.Atoi(req.FormValue("price"))
	if err != nil {
		price = 0
	}
	product := models.Product{Name: req.FormValue("name"), Price: price}

	p.db.Debug().Create(&product)

	resBytes, err := json.Marshal(product)
	if err != nil {
		res.Write([]byte(err.Error()))
		return
	}
	res.Write(resBytes)
}

func (p *Products) handlePutProduct(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		panic(err)
	}
	price, err := strconv.Atoi(req.FormValue("price"))
	if err != nil {
		price = 0
	}

	product := models.Product{}
	p.db.First(&product, id)
	product.Name = req.FormValue("name")
	product.Price = price

	p.db.Debug().Save(&product)

	resBytes, err := json.Marshal(product)
	if err != nil {
		res.Write([]byte(err.Error()))
		return
	}
	res.Write(resBytes)
}

func (p *Products) handleDelete(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		panic(err)
	}
	product := &models.Product{}
	p.db.First(product, id)
	p.db.Delete(&product)

	resBytes, err := json.Marshal(product)
	if err != nil {
		res.Write([]byte(err.Error()))
		return
	}
	res.Write(resBytes)
}
