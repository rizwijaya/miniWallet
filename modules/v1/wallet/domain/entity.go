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
type Transaction struct {
	GormModel
	WalletID    uuid.UUID `gorm:"column:wallet_id"`
	Type        int       `gorm:"column:type"`
	Amount      float64   `gorm:"column:amount"`
	ReferenceID uuid.UUID `gorm:"column:reference_id"`
	Status      int       `gorm:"column:status"`
}

type Transactions []Transaction

type ChangeStatusWalletByCustomerXID struct {
	CustomerXID uuid.UUID
	Status      int
}

type Deposit struct {
	WalletID    uuid.UUID
	Amount      float64
	ReferenceID uuid.UUID
}

type Withdrawal struct {
	WalletID    uuid.UUID
	Amount      float64
	ReferenceID uuid.UUID
}
