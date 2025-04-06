package main

import (
	"context"
	"fmt"
	"net/http"
	"slices"

	"github.com/golang-jwt/jwt/v5"
)

type Role string

const (
	Admin       Role = "Admin"
	Customer    Role = "Customer"
	Seller      Role = "Seller"
	AuthService Role = "AuthSerice"
)

type Permission int

const (
	CreateUser Permission = iota
	CreateProduct
)

var permissionTable = map[Role][]Permission{
	Admin:       {CreateUser, CreateProduct},
	Customer:    {},
	Seller:      {CreateProduct},
	AuthService: {CreateUser},
}

func HasPermission(role Role, perm Permission) bool {
	return slices.Contains(permissionTable[role], perm)
}

func verifyToken(tokenString string, secret []byte) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return secret, nil
	})

	// Check for verification errors
	if err != nil {
		return nil, err
	}

	// Check if the token is valid
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	// Return the verified token
	return token, nil
}

func (s *Server) WithAuthorizedToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenStr, err := r.Cookie("jwt")
		if err != nil {
			ResWithError(w, 401, "Unauthorized. No JWT Token.")
			return
		}

		token, err := verifyToken(tokenStr.Value, s.cfg.jwtSecret)
		if err != nil {
			ResWithError(w, 401, "Unauthorized. Invalid JWT Token.")
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			ResWithError(w, 401, "Unauthorized. Invalid token claims.")
			return
		}

		userUuid, exists := claims["user_uuid"].(string)
		if !exists || userUuid == "" {
			ResWithError(w, 401, "Unauthorized. Missing user identification.")
			return
		}

		ctx := context.WithValue(r.Context(), "user_uuid", userUuid)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
