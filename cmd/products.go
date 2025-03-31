package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	db "warehouse/internal/database"
)

func (s *Server) handleGetProduct(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	products, err := s.query.GetProduct(ctx)
	if err != nil {
		WriteError(w, 500, "Could not retrieve products")
		return
	}

	WriteJSON(w, 200, products)
}

func (s *Server) handleCreateProduct(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	fmt.Printf("req body: %v\n", r.Body)
	// Accept an array of CreateProductParams
	var reqProducts []db.CreateProductParams
	if err := json.NewDecoder(r.Body).Decode(&reqProducts); err != nil {
		WriteError(w, 400, fmt.Sprintf("malformed request body: %v", err))
		return
	}

	// Begin a transaction
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		WriteError(w, 500, fmt.Sprintf("could not begin transaction: %v", err))
		return
	}

	// Create a new querier with the transaction
	qtx := s.query.WithTx(tx)

	// Track the count of successfully created products
	insertCount := 0

	// Process each product in the array
	for _, reqP := range reqProducts {
		_, err := qtx.CreateProduct(ctx, reqP)
		if err != nil {
			// Rollback transaction on error
			tx.Rollback()
			WriteError(w, 500, fmt.Sprintf("could not insert product %d into table: %v", insertCount+1, err))
			return
		}
		insertCount++
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		WriteError(w, 500, fmt.Sprintf("could not commit transaction: %v", err))
		return
	}

	// Return only the count of created products
	response := map[string]int{"created": insertCount}
	WriteJSON(w, 200, response)
}
