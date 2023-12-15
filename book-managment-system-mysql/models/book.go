package models

import (
	"book-managment-system-mysql/config"
	"github.com/jinzhu/gorm"
	"sync"
)

var db *gorm.DB
var once sync.Once

func init() {
	once.Do(func() {
		config.Connect()
	})
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

type Book struct {
	gorm.Model
	Name        string `gorm:""json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(b)
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(ID int64) Book {
	var book Book
	db.Where("ID=?", ID).Delete(&book) // TODO: dlaczego tu jest & i czy ma być czy nie
	return book
}

//
//func init() {
//	config.Connect()
//	db = config.GetDB()
//  db.AutoMigrate(&Book{})
//}
