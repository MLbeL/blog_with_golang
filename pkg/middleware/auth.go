package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

type key string

const (
	CtxEmailKey key = "CtxEmailKey"
)

func CheckAuth(next http.Handler, secret string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		if header == "" {
			http.Error(w, "Missing Authorization: Use Your access token", http.StatusUnauthorized)
			return
		}
		if !strings.HasPrefix(header, "Bearer ") {
			http.Error(w, "Missing Authorization: Use Prefix \"Bearer {token}\"", http.StatusUnauthorized)
			return
		}
		tokenStr := strings.TrimPrefix(header, "Bearer ")
		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (any, error) {
			return []byte(secret), nil
		})
		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}
		email := token.Claims.(jwt.MapClaims)["email"]
		ctx := context.WithValue(r.Context(), CtxEmailKey, email)
		req := r.WithContext(ctx)
		next.ServeHTTP(w, req)
	})
}
