package main

import (
	"fmt"
	"net/http"
	"warehouse/internal/database"
	//db "warehouse/internal/database"
)

func (s *Server) handleGetProduct(w http.ResponseWriter, r *http.Request) {
	HandleGet(w, s.query.GetProduct, r.Context())
}

func (s *Server) handleCreateProduct(w http.ResponseWriter, r *http.Request) {
	tx, err := s.db.BeginTx(r.Context(), nil)
	if err != nil {
		ResWithError(w, 500, fmt.Sprintf("could not begin transaction: %v", err))
		return
	}
	qtx := s.query.WithTx(tx)

	HandleInsertMany(w, r, qtx.CreateProduct, tx)
}

func (s *Server) handleGetProductById(w http.ResponseWriter, r *http.Request) {
	productId, err := GetIdFromRequest(r)
	if err != nil {
		ResWithError(w, 400, fmt.Sprintf("malformed request body: %v", err))
		return
	}

	product, err := s.query.GetProductById(r.Context(), productId)
	if err != nil {
		ResWithError(w, 500, fmt.Sprintf("could not retrieve product: %v", err))
		return
	}

	allergens, err := s.query.GetAllergensForProduct(r.Context(), productId)
	if err != nil {
		ResWithError(w, 500, fmt.Sprintf("could not retrieve allergens: %v", err))
		return
	}

	res := struct {
		Product   database.Product                     `json:"product"`
		Allergens []database.GetAllergensForProductRow `json:"allergens"`
	}{
		Product:   product,
		Allergens: allergens,
	}

	ResWithJSON(w, 200, res)
}
