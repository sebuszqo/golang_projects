package handlers

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
	"simple-CRUD-API/models"
)

type MovieHandler struct {
	Movies []models.Movie
}

func NewMovieHandler(movies []models.Movie) *MovieHandler {
	return &MovieHandler{Movies: movies}
}

func (h *MovieHandler) GetMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(h.Movies)
}

func (h *MovieHandler) DeleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	initialLength := len(h.Movies)
	for index, item := range h.Movies {
		if item.ID == params["id"] {
			h.Movies = append(h.Movies[:index], h.Movies[index+1:]...)
			break
		}
	}
	if len(h.Movies) > initialLength {
		w.WriteHeader(http.StatusNotFound)
	}
	json.NewEncoder(w).Encode(h.Movies)

}

func (h *MovieHandler) GetMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range h.Movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
		}
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode("{}")
}

func (h *MovieHandler) CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie models.Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = uuid.New().String()
	h.Movies = append(h.Movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func (h *MovieHandler) UpdateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	movieID := params["id"]

	var foundMovie *models.Movie
	for _, movie := range h.Movies {
		if movie.ID == movieID {
			foundMovie = &movie
			break
		}
	}

	if foundMovie == nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"error": "Movie not found"}`))
		return
	}

	var updateMovie models.Movie
	if err := json.NewDecoder(r.Body).Decode(&updateMovie); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "Invalid JSON format"}`))
		return
	}

	if updateMovie.ISBN != "" {
		foundMovie.ISBN = updateMovie.ISBN
	}
	if updateMovie.Title != "" {
		foundMovie.Title = updateMovie.Title
	}

	if updateMovie.Director != nil {
		foundMovie.Director = updateMovie.Director
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(foundMovie)
}
