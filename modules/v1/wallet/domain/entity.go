package domain

import (
	"time"

	"github.com/google/uuid"
)

type GormModel struct {
	ID        uuid.UUID  `gorm:"column:id"`
	CreatedAt *time.Time `gorm:"created_at"`
	UpdatedAt *time.Time `gorm:"updated_at"`
}

type Wallet struct {
	GormModel
	CustomerXID uuid.UUID `gorm:"column:customer_xid"`
	Balance     float64   `gorm:"column:balance"`
	Status      int       `gorm:"column:status"`
}

type ChangeStatusWalletByCustomerXID struct {
	CustomerXID uuid.UUID
	Status      int
}
