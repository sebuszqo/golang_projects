package controllers

import (
	"book-managment-system-mysql/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

var NewBooks models.Book

func GetBook(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	newBooks := models.GetAllBooks()

	res, err := json.Marshal(newBooks)
	if err != nil {
		handleError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {

}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		handleError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	newBook, db := models.GetBookById(ID)
	if db.Error != nil {
		handleError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	res, err := json.Marshal(newBook)
	if err != nil {
		handleError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func UpdateBook(w http.ResponseWriter, r *http.Request) {

}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		handleError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	newBook := models.DeleteBook(ID)
	res, err := json.Marshal(newBook)
	if err != nil {
		handleError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func handleError(w http.ResponseWriter, status int, errorMessage string) {
	w.WriteHeader(status)
	w.Write([]byte(`{"error": "` + errorMessage + `"}`))
}
