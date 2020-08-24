package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/k8shop/go-rest-api/pkg/handlers"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := initDB()
	defer db.Close()

	if err != nil {
		panic(err)
	}
	router := mux.NewRouter()
	err = handlers.Register(router, db)
	if err != nil {
		panic(err)
	}

	http.ListenAndServe(":8080", router)
}

func initDB() (*sql.DB, error) {
	connectionURL := "root:th00perThecure@tcp(database:3306)/"
	log.Printf("Connecting to DB: %s", connectionURL)
	db, err := sql.Open("mysql", connectionURL)
	if err != nil {
		return nil, err
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS bikepacker")
	if err != nil {
		return nil, err
	}

	_, err = db.Exec("USE bikepacker")
	if err != nil {
		return nil, err
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS bikes (id INT NOT NULL AUTO_INCREMENT, name VARCHAR(256) NOT NULL, PRIMARY KEY (id))")
	if err != nil {
		return nil, err
	}

	return db, nil
}
