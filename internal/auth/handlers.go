package auth

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/yogtanko/go-kios/internal/json"
	"github.com/yogtanko/go-kios/internal/middleware"
)

type handler struct {
	jwtSecret []byte
}

func NewHandler(secret []byte) *handler {
	return &handler{
		jwtSecret: secret,
	}
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *handler) Login(w http.ResponseWriter, r *http.Request) {
	req, err := json.Read[LoginRequest](r)
	if err != nil {
		slog.Error("invalid request body", "error", err)
		json.Write(w, http.StatusBadRequest, map[string]string{"message": "invalid request body"})
		return
	}

	if req.Username != "admin" || req.Password != "password" {
		slog.Error("invalid credentials", "error", err)
		json.Write(w, http.StatusUnauthorized, map[string]string{"message": "invalid credentials"})
		return
	}

	claims := &middleware.Claims{
		UserName: req.Username,
		Roles:    []string{"customer"},
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "go-kios",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(h.jwtSecret)
	if err != nil {
		slog.Error("failed to generate token", "error", err)
		json.Write(w, http.StatusInternalServerError, map[string]string{"message": "failed to generate token"})
		return
	}
	slog.Info("Authorization success", "token", tokenStr)
	json.Write(w, http.StatusOK, map[string]string{"token": tokenStr})
}
