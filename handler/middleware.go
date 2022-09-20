package handler

import (
	"net/http"

	"github.com/y-magavel/go-todo-api/auth"
)

func AuthMiddleware(j *auth.JWTer) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			req, err := j.FillContext(r)
			if err != nil {
				RespondJSON(r.Context(), w, ErrResponse{Message: "not find auth info", Details: []string{err.Error()}}, http.StatusUnauthorized)
				return
			}
			next.ServeHTTP(w, req)
		})
	}
}
