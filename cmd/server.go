package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"warehouse/internal/database"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Server struct {
	router *chi.Mux
	query  *database.Queries
	db     *sql.DB
}

func NewServer(db *sql.DB, query *database.Queries) Server {

	r := chi.NewRouter()

	s := Server{
		router: r,
		query:  query,
		db:     db,
	}

	r.Use(middleware.Logger)

	r.Get("/api/product", s.handleGetProduct)
	r.Post("/api/product", s.handleCreateProduct)
	r.Get("/api/product/{id}", s.handleGetProductById)

	r.Get("/api/allergen", s.handleGetAllergen)
	r.Post("/api/allergen", s.handleCreateAllergen)
	r.Get("/api/allergen/{id}", s.handleGetAllergenById)

	r.Get("/api/user/{uuid}", s.handleGetUser)
	r.Post("/api/user", s.handleCreateUser)

	return s
}

func (s *Server) Run(port string) {
	log.Printf("Server listening on port %s...\n", port)
	err := http.ListenAndServe(":"+port, s.router)
	if err != nil {
		log.Fatal(err)
	}
}

func WriteJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Add("Content-Type", "application/json")

	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marashal respone payload! %v\n", err)
		WriteError(w, 500, "Failed to marshal payload")
	}

	w.WriteHeader(status)
	w.Write(data)
}

func WriteError(w http.ResponseWriter, status int, msg string) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	type Res struct {
		Error string `json:"error"`
	}

	json.NewEncoder(w).Encode(Res{Error: msg})

	if status > 499 {
		log.Printf("RESPONSE SATUS 5xx: %s", msg)
	}
}

func GetIdFromRequest(r *http.Request) (int64, error) {
	id := chi.URLParam(r, "id")
	if id == "" {
		return 0, fmt.Errorf("empty ID")
	}

	return strconv.ParseInt(id, 10, 64)
}
