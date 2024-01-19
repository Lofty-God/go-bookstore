package main

import (
	"log"
	"net/http"

	"github.com/Lofty-God/go-bookstore/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	http.Handle("/", r)
	routes.RegisterBookStoreRoutes(r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
