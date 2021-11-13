package http

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// Handler -
type Handler struct {
	Router *mux.Router
}

// Response objecgi
type Response struct {
	Message string
	Error   string
}

// ErrorResponse -
type ErrorResponse struct {
	Error string
}

// New - returns a new handler
func New() *Handler {
	return &Handler{}
}

// SetupRoutes sets up all the routes for the app
func (h *Handler) SetupRoutes() {
	h.Router = mux.NewRouter()

	h.Router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(Response{Message: "I am Alive!"}); err != nil {
			panic(err)
		}
	})

	h.Router.HandleFunc("/v1/execute", h.ExecuteChallenge).Methods("POST")
}

func sendErrorResponse(w http.ResponseWriter, message string, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	if err := json.NewEncoder(w).Encode(Response{Message: message, Error: err.Error()}); err != nil {
		panic(err)
	}
}
