package repository

import (
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/google/uuid"
	"github.com/rizwijaya/miniWallet/modules/v1/wallet/domain"
	"gorm.io/gorm"
)

type WalletRepository interface {
	GetWalletByUserID(userID uuid.UUID) (domain.Wallet, error)
	CreateWallet(wallet domain.Wallet) error
}

type walletRepository struct {
	db *gorm.DB
	mc *memcache.Client
}

func NewRepository(db *gorm.DB, mc *memcache.Client) *walletRepository {
	return &walletRepository{
		db: db,
		mc: mc,
	}
}
