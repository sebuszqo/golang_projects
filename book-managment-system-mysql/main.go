package main

import (
	"book-managment-system-mysql/routes"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Println("Starting server at: http://localhost:8080")
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
