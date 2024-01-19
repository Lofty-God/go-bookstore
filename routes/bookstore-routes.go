package routes

import (
	"github.com/Lofty-God/go-bookstore/controllers"

	"github.com/gorilla/mux"
)

func RegisterBookStoreRoutes(r *mux.Router) {

	r.HandleFunc("/book/", controllers.GetAllBooks).Methods("GET")
	r.HandleFunc("/book/{id}", controllers.GetBookById).Methods("GET")
	r.HandleFunc("/books/", controllers.CreateBook).Methods("POST")
	r.HandleFunc("/books/{id}", controllers.UpdateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", controllers.DeleteBook).Methods("DELETE")
}
