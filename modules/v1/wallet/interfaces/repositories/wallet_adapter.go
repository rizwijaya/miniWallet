package repository

import (
	"github.com/google/uuid"
	"github.com/rizwijaya/miniWallet/modules/v1/wallet/domain"
	"gorm.io/gorm"
)

type WalletRepository interface {
	GetWalletByCustomerXID(customerXID uuid.UUID) (domain.Wallet, error)
	GetWalletByID(walletID uuid.UUID) (domain.Wallet, error)
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
}

func NewRepository(db *gorm.DB) *walletRepository {
	return &walletRepository{
		db: db,
	}
}
