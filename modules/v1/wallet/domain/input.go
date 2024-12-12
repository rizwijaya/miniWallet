package domain

import "github.com/google/uuid"

type InitMyAccountInput struct {
	CustomerXID uuid.UUID `form:"customer_xid"`
}

type DepositInput struct {
	Amount      float64   `form:"amount"`
	ReferenceID uuid.UUID `form:"reference_id"`
}

type WithdrawalInput struct {
	Amount      float64   `form:"amount"`
	ReferenceID uuid.UUID `form:"reference_id"`
}
