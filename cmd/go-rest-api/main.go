package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/k8shop/go-rest-api/pkg/handlers"
	"github.com/k8shop/go-rest-api/pkg/handlers/middleware"

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
	router.Use(middleware.AddCommonHeaders)
	err = handlers.Register(router, db)
	if err != nil {
		panic(err)
	}

	http.ListenAndServe(":8080", router)
}

func initDB() (*sql.DB, error) {
	connectionURL := fmt.Sprintf(
		"%v:%v@tcp(%v)/",
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_PASS"),
		os.Getenv("DATABASE_HOST"),
	)
	log.Printf("Connecting to DB: %s", connectionURL)
	db, err := sql.Open("mysql", connectionURL)
	if err != nil {
		return nil, err
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + os.Getenv("DATABASE_NAME"))
	if err != nil {
		return nil, err
	}

	connectionURL = fmt.Sprintf(
		"%v:%v@tcp(%v)/%v",
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_PASS"),
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_NAME"),
	)
	log.Printf("Connecting to DB: %s", connectionURL)
	db, err = sql.Open("mysql", connectionURL)
	if err != nil {
		return nil, err
	}

	schemaDir, err := os.Open("./schema")
	if err != nil {
		return nil, err
	}
	files, err := schemaDir.Readdirnames(-1)
	if err != nil {
		return nil, err
	}

	for _, fileName := range files {
		if !strings.Contains(fileName, ".tbl.sql") {
			continue
		}
		buf, err := ioutil.ReadFile("./schema/" + fileName)
		if err != nil {
			return nil, err
		}
		_, err = db.Exec(string(buf))
		if err != nil {
			return nil, err
		}
	}

	return db, nil
}
