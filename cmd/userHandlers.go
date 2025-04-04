package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"warehouse/internal/database"

	"github.com/go-chi/chi/v5"
)

func (s *Server) handleGetUser(w http.ResponseWriter, r *http.Request) {
	uuid := chi.URLParam(r, "uuid")

	user, err := s.query.GetUser(r.Context(), uuid)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			WriteError(w, 400, "No such user with uuid: "+uuid)
		} else {
			WriteError(w, 500, fmt.Sprintf("could retrieve user from database: %v", err))
		}
		return
	}

	WriteJSON(w, 200, user)
}

func (s *Server) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	var payload database.CreateUserParams
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		WriteError(w, 400, fmt.Sprintf("malformed request body: %v", err))
		return
	}

	user, err := s.query.CreateUser(r.Context(), payload)
	if err != nil {
		WriteError(w, 500, fmt.Sprintf("could not save user in database, %v", err))
	}

	WriteJSON(w, 200, "Successfully saved user "+user.Uuid)
}
