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
			ID:    "f47ac10b-58cc-0372-8567-0e02b2c3d479",
			ISBN:  "435553",
			Title: "Oppenheimer",
			Director: &models.Director{
				FirstName: "Christopher",
				LastName:  "Nolan",
			},
		},
		{
			ID:    "1e2d5a26-51a0-2d1e-1886-970b9b9bae90",
			ISBN:  "342144",
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
	r.HandleFunc("/movies/{id}", movieHandler.UpdateMovie).Methods(http.MethodPatch)
	r.HandleFunc("/movies/{id}", movieHandler.DeleteMovie).Methods(http.MethodDelete)

	fmt.Printf("Starting server at: %s\n", "localhost:8080")

	log.Fatal(http.ListenAndServe(":8080", r))
}
