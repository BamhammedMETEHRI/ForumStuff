package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/BamhammedMETEHRI/go-bookstore/pkg/models"
	"github.com/BamhammedMETEHRI/go-bookstore/pkg/utils"
	"github.com/gorilla/mux"
)

var NewBook models.Book

func GetBook(rw http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	rw.Header().Set("Content-Type", "pkglication/json")
	rw.WriteHeader(http.StatusOK)
	rw.Write(res)
}

func GetBookById(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	Id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookDetails, _ := models.GetBookById(Id)
	res, _ := json.Marshal(bookDetails)
	rw.Header().Set("Content-Type", "pkglication/json")
	rw.WriteHeader(http.StatusOK)
	rw.Write(res)
}

func CreateBook(rw http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	b := CreateBook.CreateBook()
	res, _ := json.Marshal(b)
	rw.Header().Set("Content-Type", "pkglication/json")
	rw.WriteHeader(http.StatusOK)
	rw.Write(res)
}

func DeleteBook(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	Id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing", err)
	}
	book := models.DeleteBook(Id)
	res, _ := json.Marshal(book)
	rw.Header().Set("Content-Type", "pkglication/json")
	rw.WriteHeader(http.StatusOK)
	rw.Write(res)
}

func UpdateBook(rw http.ResponseWriter, r *http.Request) {
	var updateBook = &models.Book{}
	utils.ParseBody(r, updateBook)
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	Id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing", err)
	}
	booksDetails, db := models.GetBookById(Id)
	if updateBook.Name != "" {
		booksDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		booksDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		booksDetails.Publication = updateBook.Publication
	}
	db.Save(&booksDetails)
	res, _ := json.Marshal(booksDetails)
	rw.Header().Set("Content-Type", "pkglication/json")
	rw.WriteHeader(http.StatusOK)
	rw.Write(res)
}
