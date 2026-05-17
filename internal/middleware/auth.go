package middleware

import (
	"context"
	"log/slog"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/yogtanko/go-kios/internal/json"
)

type contextKey string

var jwtSecret []byte

func Init(secret []byte) {
	jwtSecret = secret
}

const ClaimsKey contextKey = "claims"

type Claims struct {
	Roles    []string `json:"roles"`
	UserName string   `json:"username"`
	jwt.RegisteredClaims
}

func Authenticator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			slog.Error("missing authorization header", "error", "authHeader == \"\"")
			json.Write(w, http.StatusUnauthorized, map[string]string{"message": "missing authorization header"})
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			slog.Error("invalid authorization format", "error", "len(parts) != 2 || strings.ToLower(parts[0]) != \"bearer\"")
			json.Write(w, http.StatusUnauthorized, map[string]string{"message": "invalid authorization format"})
			return
		}

		tokenStr := parts[1]

		claims := &Claims{}

		token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (any, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return jwtSecret, nil
		})
		if err != nil || !token.Valid {
			slog.Error("invalid or expired token", "error", err.Error())
			json.Write(w, http.StatusUnauthorized, map[string]string{"message": "invalid or expired token"})
			return
		}
		ctx := context.WithValue(r.Context(), ClaimsKey, claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetClaims(r *http.Request) *Claims {
	claims, _ := r.Context().Value(ClaimsKey).(*Claims)
	return claims
}
