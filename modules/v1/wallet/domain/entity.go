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
	UserID  uuid.UUID `gorm:"column:user_id"`
	Balance float64   `gorm:"column:balance"`
	Status  int       `gorm:"column:status"`
}

type ChangeStatusWalletByUserID struct {
	UserID uuid.UUID
	Status int
}
