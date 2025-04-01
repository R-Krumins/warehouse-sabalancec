package main

import (
	"context"
	"net/http"
)

func DBGet[T any](w http.ResponseWriter, f func(context.Context) ([]T, error), ctx context.Context) {
	data, err := f(ctx)
	if err != nil {
		WriteError(w, 500, "Could not retrieve data from database")
		return
	}

	WriteJSON(w, 200, data)
}

func DBInsert[T any, P any](ctx context.Context, payload P, insertFunc func(context.Context, P) (T, error)) (T, error) {
	return insertFunc(ctx, payload)
}
