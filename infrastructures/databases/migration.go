package database

import (
	"time"

	"github.com/google/uuid"
)

type GormModel struct {
	ID        uuid.UUID  `gorm:"column:id;type:uuid;primaryKey;not null"`
	CreatedAt *time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt *time.Time `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP;autoUpdateTime"`
}

type Wallet struct {
	GormModel
	CustomerXID  uuid.UUID     `gorm:"column:customer_xid;type:uuid;not null"`
	Balance      float64       `gorm:"column:balance;type:decimal(20,2)"`
	Status       int           `gorm:"column:status;type:int;default:0"`
	Transactions []Transaction `gorm:"foreignKey:WalletID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"` // One-to-Many
}

type Transaction struct {
	GormModel
	WalletID    uuid.UUID `gorm:"column:wallet_id"`
	Type        int       `gorm:"column:type"`
	Amount      float64   `gorm:"column:amount"`
	ReferenceID uuid.UUID `gorm:"column:reference_id;type:uuid;unique;not null"`
	Status      int       `gorm:"column:status;type:int;default:0"`
}
