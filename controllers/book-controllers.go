package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/Lofty-God/go-bookstore/models"
	"github.com/Lofty-God/go-bookstore/utils"
)

var newBook models.Book

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	getNewBooks := models.GetAllBooks()
	res, _ := json.Marshal(getNewBooks)
	w.Header().Set("content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	bookId := vars["bookId"]
	Id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing data")
	}
	BookDetails, _ := models.GetBookById(Id)
	res, _ := json.Marshal(BookDetails)
	w.Header().Set("content-type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	createbook := models.Book{}
	utils.ParseBody(r, createbook)
	b := createbook.CreateBook()
	res, _ := json.Marshal(b)
	w.Header().Set("content-type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	Id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Printf("error occurred while parsing")
	}
	deletebook := models.DeleteBookById(Id)
	res, _ := json.Marshal(deletebook)
	w.Header().Set("content-type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var updateBook = &models.Book{}
	utils.ParseBody(r, updateBook)
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	Id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookDetails, db := models.GetBookById(Id)
	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Title != "" {
		bookDetails.Title = updateBook.Title
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("content-type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
	db.Save(bookDetails)

}
