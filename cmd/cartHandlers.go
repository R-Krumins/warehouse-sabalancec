package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"warehouse/internal/database"
)

func (s *Server) handleAddToCart(w http.ResponseWriter, r *http.Request) {
	var payload database.PatchCartParams
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		ResWithError(w, 400, fmt.Sprintf("Malformed request body: %v", err))
		return
	}

	item, err := s.query.PatchCart(r.Context(), payload)
	if err != nil {
		ResWithError(w, 500, fmt.Sprintf("Internal Server Error: %v", err))
		return
	}

	ResWithJSON(w, 200, item)
}

func (s *Server) handleGetCartForUser(w http.ResponseWriter, r *http.Request) {
	userUuid, _ := r.Cookie("user_uuid")

	items, err := s.query.GetCartForUser(r.Context(), userUuid.Value)
	if err != nil {
		ResWithError(w, 500, fmt.Sprintf("Internal Server Error: %v", err))
		return
	}

	ResWithJSON(w, 200, items)
}
