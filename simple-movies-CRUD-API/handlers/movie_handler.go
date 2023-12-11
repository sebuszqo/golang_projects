package handlers

import (
	"encoding/json"
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
