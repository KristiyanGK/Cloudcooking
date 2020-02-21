package middlewares

import (
	"strings"
	"context"
	"github.com/KristiyanGK/cloudcooking/api/auth"
	"net/http"
)

// AuthenticationMiddleware checks wheater the user is authenticated via jwt token
// if the token is missing or is invalid a 401 unathorized is sent to the client
// if the token is valid loads the token into context
func AuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		const BearerScheme = "Bearer"

		tokenHeader := r.Header.Get("Authorization")

		tokens := strings.Split(tokenHeader, " ")

		if len(tokens) != 2 || tokens[0] != BearerScheme || auth.IsTokenValid(tokens[1]) {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "token", tokens[1])
		
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}