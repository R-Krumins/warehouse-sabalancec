package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"warehouse/internal/database"

	"github.com/go-chi/chi/v5"
)

func (s *Server) CartRouter(r chi.Router) {
	r.Use(s.WithAuthorizedToken)
	r.Get("/", s.handleGetCartForUser)
	r.Patch("/", s.handleAddToCart)
}

func (s *Server) handleAddToCart(w http.ResponseWriter, r *http.Request) {
	var newItem database.PatchCartParams
	if err := json.NewDecoder(r.Body).Decode(&newItem); err != nil {
		ResWithError(w, 400, fmt.Sprintf("Malformed request body: %v", err))
		return
	}

	newItem.UserUuid = r.Context().Value("user_uuid").(string)

	item, err := s.query.PatchCart(r.Context(), newItem)
	if err != nil {
		ResWithError(w, 500, fmt.Sprintf("Internal Server Error: %v", err))
		return
	}

	ResWithJSON(w, 200, item)
}

func (s *Server) handleGetCartForUser(w http.ResponseWriter, r *http.Request) {
	userUuid := r.Context().Value("user_uuid").(string)

	items, err := s.query.GetCartForUser(r.Context(), userUuid)
	if err != nil {
		ResWithError(w, 500, fmt.Sprintf("Internal Server Error: %v", err))
		return
	}

	ResWithJSON(w, 200, items)
}
