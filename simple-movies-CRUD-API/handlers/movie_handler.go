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
