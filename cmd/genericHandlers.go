package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

type GetDBFunc[M any] func(context.Context) ([]M, error)

// types: [M]odel
func HandleGet[M any](w http.ResponseWriter, dbFunc GetDBFunc[M], ctx context.Context) {
	data, err := dbFunc(ctx)
	if err != nil {
		WriteError(w, 500, "Could not retrieve data from database")
		return
	}

	WriteJSON(w, 200, data)
}

type InsertDBFunc[M any, P any] func(context.Context, P) (M, error)

// types: [M]odel, [P]ayload
func HandleInsertMany[M any, P any](w http.ResponseWriter, r *http.Request, dbFunc InsertDBFunc[M, P], tx *sql.Tx) {
	var payload []P
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		WriteError(w, 400, fmt.Sprintf("malformed request body: %v", err))
		return
	}

	insertCount := 0
	for _, p := range payload {
		_, err := dbFunc(r.Context(), p)
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

type GetByIdDBFunc[M any] func(context.Context, int64) (M, error)

// types: [M]odel
func HandleGetById[M any](w http.ResponseWriter, r *http.Request, dbFunc GetByIdDBFunc[M]) {
	id, err := GetIdFromRequest(r)
	if err != nil {
		WriteError(w, 400, "malformed request: id must be an integer")
		return
	}

	data, err := dbFunc(r.Context(), id)
	if err != nil {
		WriteError(w, 500, fmt.Sprintf("could not retrieve data from database: %v", err))
		return
	}

	WriteJSON(w, 200, data)
}
