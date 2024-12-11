package domain

import (
	"time"

	"github.com/google/uuid"
)

type WalletResponse struct {
	ID        uuid.UUID `json:"id"`
	OwnedBy   uuid.UUID `json:"owned_by"`
	Status    string    `json:"status"`
	EnabledAt time.Time `json:"enabled_at"`
	Balance   float64   `json:"balance"`
}
