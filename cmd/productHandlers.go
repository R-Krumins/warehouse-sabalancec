package main

import (
	"fmt"
	"net/http"
	//db "warehouse/internal/database"
)

func (s *Server) handleGetProduct(w http.ResponseWriter, r *http.Request) {
	HandleGet(w, s.query.GetProduct, r.Context())
}

func (s *Server) handleCreateProduct(w http.ResponseWriter, r *http.Request) {
	tx, err := s.db.BeginTx(r.Context(), nil)
	if err != nil {
		WriteError(w, 500, fmt.Sprintf("could not begin transaction: %v", err))
		return
	}
	qtx := s.query.WithTx(tx)

	HandleInsertMany(w, r, qtx.CreateProduct, tx)
}

func (s *Server) handleGetProductById(w http.ResponseWriter, r *http.Request) {
	HandleGetById(w, r, s.query.GetProductById)
}
