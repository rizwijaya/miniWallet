package domain

import "github.com/google/uuid"

type InitMyAccountInput struct {
	CustomerXID uuid.UUID `form:"customer_xid"`
}
