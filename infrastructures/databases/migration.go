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
	UserID       uuid.UUID     `gorm:"column:user_id;type:uuid;not null"`
	Balance      float64       `gorm:"column:balance;type:decimal(20,2)"`
	Status       int           `gorm:"column:status;type:int;default:0"`
	Transactions []Transaction `gorm:"foreignKey:WalletID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"` // One-to-Many
}

type Transaction struct {
	GormModel
	WalletID    uuid.UUID `gorm:"column:wallet_id;type:uuid;not null;index"`
	Type        int       `gorm:"column:type;type:int;not null"`
	Amount      float64   `gorm:"column:amount;type:decimal(20,2);not null"`
	ReferenceID string    `gorm:"column:reference_id;type:varchar(255);unique;not null"`
	Status      int       `gorm:"column:status;type:int;default:0"`
}
