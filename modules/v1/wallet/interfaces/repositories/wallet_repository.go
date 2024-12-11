package repository

import (
	"github.com/bradfitz/gomemcache/memcache"
	"gorm.io/gorm"
)

type WalletRepository interface {
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
