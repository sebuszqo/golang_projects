package controllers

import (
	"book-managment-system-mysql/models"
	"book-managment-system-mysql/utils"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

var NewBooks models.Book

func IsMethodAllowed(w http.ResponseWriter, r *http.Request, allowedMethod string) bool {
	if r.Method != allowedMethod {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return false
	}
	return true
}

func handleError(w http.ResponseWriter, status int, errorMessage string) {
	w.WriteHeader(status)
	w.Write([]byte(`{"error": "` + errorMessage + `"}`))
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	if !IsMethodAllowed(w, r, http.MethodGet) {
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
	if !IsMethodAllowed(w, r, http.MethodPost) {
		return
	}

	CreateBook := &models.Book{}
	if err := utils.ParseBody(r, CreateBook); err != nil {
		handleError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	b := CreateBook.CreateBook()
	res, err := json.Marshal(b)
	if err != nil {
		handleError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	if !IsMethodAllowed(w, r, http.MethodGet) {
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
	if !IsMethodAllowed(w, r, http.MethodPut) {
		return
	}

	var updateBook = &models.Book{}
	if err := utils.ParseBody(r, updateBook); err != nil {
		handleError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		handleError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	bookDetails, db := models.GetBookById(ID)
	if db.Error != nil {
		handleError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}
	db.Save(&bookDetails)

	res, err := json.Marshal(bookDetails)
	if err != nil {
		handleError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	if !IsMethodAllowed(w, r, http.MethodDelete) {
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
