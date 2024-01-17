package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func registerBookStoreRoutes(r *mux.Router) {
	http.Handle("/", r)
	r.HandleFunc("/book/", controllers.getAllBook).Methods("GET")
	r.HandleFunc("/book/{id}", controllers.getBookById).Methods("GET")
	r.HandleFunc("/books/", controllers.createBook).Methods("POST")
	r.HandleFunc("/books/{id}", controllers.updateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", controllers.deleteBook).Methods("DELETE")
}
