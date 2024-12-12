package repository

import (
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/google/uuid"
	"github.com/rizwijaya/miniWallet/modules/v1/wallet/domain"
	"gorm.io/gorm"
)

type WalletRepository interface {
	GetWalletByCustomerXID(CustGetWalletByCustomerXID uuid.UUID) (domain.Wallet, error)
	CreateWallet(wallet domain.Wallet) error
	ChangeStatusWalletByCustomerXID(param domain.ChangeStatusWalletByCustomerXID) (domain.Wallet, error)
	GetTransactionsByCustomerXID(customerXID uuid.UUID) (domain.Transactions, error)
	GetDBTx() *gorm.DB
	GetWalletByIDWithTx(tx *gorm.DB, walletID uuid.UUID) (domain.Wallet, error)
	UpdateWalletWithTx(tx *gorm.DB, wallet domain.Wallet) error
	CreateTransaction(transaction domain.Transaction) error
	GetTransactionsByReferenceID(referenceID uuid.UUID) (domain.Transactions, error)
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
