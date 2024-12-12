package domain

import (
	"time"

	"github.com/google/uuid"
)

type WalletEnableResponse struct {
	ID        uuid.UUID `json:"id"`
	OwnedBy   uuid.UUID `json:"owned_by"`
	Status    string    `json:"status"`
	EnabledAt time.Time `json:"enabled_at"`
	Balance   float64   `json:"balance"`
}

type WalletDisableResponse struct {
	ID         uuid.UUID `json:"id"`
	OwnedBy    uuid.UUID `json:"owned_by"`
	Status     string    `json:"status"`
	DisabledAt time.Time `json:"disabled_at"`
	Balance    float64   `json:"balance"`
}

type TransactionResponse struct {
	ID           uuid.UUID `json:"id"`
	Status       string    `json:"status"`
	TransactedAt time.Time `json:"transacted_at"`
	Type         string    `json:"type"`
	Amount       float64   `json:"amount"`
	ReferenceID  uuid.UUID `json:"reference_id"`
}

type DepositResponse struct {
	ID          uuid.UUID `json:"id"`
	DepositedBy uuid.UUID `json:"deposited_by"`
	Status      string    `json:"status"`
	DepositedAt time.Time `json:"deposited_at"`
	Amount      float64   `json:"amount"`
	ReferenceID uuid.UUID `json:"reference_id"`
}

type WithdrawalResponse struct {
	ID           uuid.UUID `json:"id"`
	WithdrawalBy uuid.UUID `json:"withdrawn_by"`
	Status       string    `json:"status"`
	WithdrawalAt time.Time `json:"withdrawn_at"`
	Amount       float64   `json:"amount"`
	ReferenceID  uuid.UUID `json:"reference_id"`
}
