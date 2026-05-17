package products

import (
	"log/slog"
	"net/http"

	"github.com/yogtanko/go-kios/internal/json"
)

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) ListProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.service.ListProducts(r.Context(), r.RemoteAddr)
	if err != nil {
		slog.Error("Terjadi Kesalahan", "error", err)
		json.Write(w, http.StatusInternalServerError, map[string]string{"message": "Terjadi Kesalahan"})
		return
	}
	slog.Info("ListProducts", "data", products)
	json.Write(w, http.StatusOK, products)
}
