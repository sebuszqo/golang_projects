package handlers

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
	"simple-CRUD-API/models"
)

// MovieHandler handles HTTP requests related to movies.
type MovieHandler struct {
	Movies []models.Movie
}

// NewMovieHandler creates a new instance of MovieHandler.
func NewMovieHandler(movies []models.Movie) *MovieHandler {
	return &MovieHandler{Movies: movies}
}

// GetMovies handles the HTTP GET request to retrieve all movies.
func (h *MovieHandler) GetMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(h.Movies)
}

// DeleteMovie handles the HTTP DELETE request to delete a movie by ID.
func (h *MovieHandler) DeleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	initialLength := len(h.Movies)

	// Find the movie by ID and remove it from the slice.
	for index, item := range h.Movies {
		if item.ID == params["id"] {
			h.Movies = append(h.Movies[:index], h.Movies[index+1:]...)
			break
		}
	}

	// Check if the movie was found and deleted.
	if len(h.Movies) == initialLength {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"error": "Movie not found"}`))
		return
	}

	// Respond with the updated list of movies.
	json.NewEncoder(w).Encode(h.Movies)
}

// GetMovie handles the HTTP GET request to retrieve a movie by ID.
func (h *MovieHandler) GetMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	// Find the movie by ID and respond with its details.
	for _, item := range h.Movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	// Respond with an error if the movie is not found.
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"error": "Movie not found"}`))
}

// CreateMovie handles the HTTP POST request to create a new movie.
func (h *MovieHandler) CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie models.Movie

	// Decode the incoming JSON request to create a new movie.
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = uuid.New().String()
	h.Movies = append(h.Movies, movie)

	// Respond with the details of the newly created movie.
	json.NewEncoder(w).Encode(movie)
}

// UpdateMovie handles the HTTP PUT request to update a movie by ID.
func (h *MovieHandler) UpdateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	movieID := params["id"]

	// Find the movie by ID in the slice.
	var foundMovie *models.Movie
	for i, movie := range h.Movies {
		if movie.ID == movieID {
			foundMovie = &h.Movies[i]
			break
		}
	}

	// Respond with an error if the movie is not found.
	if foundMovie == nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"error": "Movie not found"}`))
		return
	}

	// Decode the incoming JSON request to update the movie.
	var updateMovie models.Movie
	if err := json.NewDecoder(r.Body).Decode(&updateMovie); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "Invalid JSON format"}`))
		return
	}

	// Update the movie fields if provided in the request.
	if updateMovie.ISBN != "" {
		foundMovie.ISBN = updateMovie.ISBN
	}
	if updateMovie.Title != "" {
		foundMovie.Title = updateMovie.Title
	}
	if updateMovie.Director != nil {
		foundMovie.Director = updateMovie.Director
	}

	// Respond with the updated movie details.
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(foundMovie)
}
