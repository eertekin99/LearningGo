package main

import (
	"log"
	"net/http"

	//absolute path
	"github.com/eertekin99/go-bookstore/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}

//_ "github.com/jinzhu/gorm/dialects/mysql"
//_ "github.com/jinzhu/gorm/dialects/sqlite"
//_ "github.com/mattn/go-sqlite3"
