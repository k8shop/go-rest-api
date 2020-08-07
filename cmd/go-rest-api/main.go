package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/k8shop/go-rest-api/pkg/handlers"
)

func main() {
	router := mux.NewRouter()
	err := handlers.Register(router)
	if err != nil {
		log.Panic(err)
	}

	http.ListenAndServe(":8080", router)
}
