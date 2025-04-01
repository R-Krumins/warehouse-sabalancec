package main

import (
	"net/http"
)

func (s *Server) handleGetAllergen(w http.ResponseWriter, r *http.Request) {
	HandleGet(w, s.query.GetAllergen, r.Context())
}

func (s *Server) handleCreateAllergen(w http.ResponseWriter, r *http.Request) {
	tx, err := s.db.BeginTx(r.Context(), nil)
	if err != nil {
		WriteError(w, 500, "could not begin transaction")
		return
	}
	qtx := s.query.WithTx(tx)

	HandleInsertMany(w, r, qtx.CreateAllergen, tx)
}

func (s *Server) handleGetAllergenById(w http.ResponseWriter, r *http.Request) {
	HandleGetById(w, r, s.query.GetAllergenById)
}
