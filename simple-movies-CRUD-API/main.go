package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"simple-CRUD-API/handlers"
	"simple-CRUD-API/models"
)

type Director struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func main() {
	r := mux.NewRouter()

	movies := []models.Movie{
		{
			Id:    "1",
			Isbn:  "435553",
			Title: "Oppenheimer",
			Director: &models.Director{
				FirstName: "Christopher",
				LastName:  "Nolan",
			},
		},
		{
			Id:    "2",
			Isbn:  "342144",
			Title: "Interstellar",
			Director: &models.Director{
				FirstName: "Christopher",
				LastName:  "Nolan",
			},
		},
	}

	movieHandler := handlers.NewMovieHandler(movies)
	r.HandleFunc("/movies", movieHandler.GetMovies).Methods(http.MethodGet)
	r.HandleFunc("/movies/{id}", movieHandler.GetMovie).Methods(http.MethodGet)
	r.HandleFunc("/movies", movieHandler.CreateMovie).Methods(http.MethodPost)
	r.HandleFunc("/movies/{id}", movieHandler.UpdateMovie).Methods(http.MethodPut)
	r.HandleFunc("/movies/{id}", movieHandler.DeleteMovie).Methods(http.MethodDelete)

	fmt.Printf("Starting server at: %s\n", "localhost:8080")

	log.Fatal(http.ListenAndServe(":8080", r))
}
