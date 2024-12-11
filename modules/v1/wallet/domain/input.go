package domain

import "github.com/google/uuid"

type InitMyAccountInput struct {
	UserID uuid.UUID `form:"customer_xid"`
}
