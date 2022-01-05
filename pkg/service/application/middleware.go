package application

import (
	"context"
	"net/http"
)

type UserIDKey string

const UIDK UserIDKey = "user-id"

func StubAuthMiddleware(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		if id == "" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("missing id parameter in request"))
			return
		}

		ctx := context.WithValue(r.Context(), UIDK, id)
		f(w, r.WithContext(ctx))
	}
}
