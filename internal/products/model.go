package products

import (
	"time"

	"github.com/google/uuid"
)

type Products struct {
	Id        uuid.UUID `json:"id" db:"id"`
	Code      string    `json:"code" db:"code"`
	Name      string    `json:"name" db:"name"`
	Barcode   string    `json:"barcode" db:"barcode"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
