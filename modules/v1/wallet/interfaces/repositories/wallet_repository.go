package repository

import (
	"github.com/google/uuid"
	"github.com/rizwijaya/miniWallet/modules/v1/wallet/domain"
	"gorm.io/gorm"
)

func (wr *walletRepository) GetWalletByCustomerXID(customerXID uuid.UUID) (domain.Wallet, error) {
	var wallet domain.Wallet
	if err := wr.db.Model(&domain.Wallet{}).Where("customer_xid = ?", customerXID).First(&wallet).Error; err != nil {
		return domain.Wallet{}, err
	}
	return wallet, nil
}

func (wr *walletRepository) CreateWallet(wallet domain.Wallet) error {
	return wr.db.Create(&wallet).Error
}

func (wr *walletRepository) ChangeStatusWalletByCustomerXID(param domain.ChangeStatusWalletByCustomerXID) (domain.Wallet, error) {
	var wallet domain.Wallet
	if err := wr.db.Model(&domain.Wallet{}).Where("customer_xid = ?", param.CustomerXID).Update("status", param.Status).Scan(&wallet).Error; err != nil {
		return domain.Wallet{}, err
	}

	return wallet, nil
}

func (wr *walletRepository) GetTransactionsByCustomerXID(customerXID uuid.UUID) (domain.Transactions, error) {
	var transactions domain.Transactions
	if err := wr.db.Find(&transactions).Error; err != nil {
		return nil, err
	}

	return transactions, nil
}

func (wr *walletRepository) GetDBTx() *gorm.DB {
	return wr.db.Begin()
}

func (wr *walletRepository) GetWalletByIDWithTx(tx *gorm.DB, walletID uuid.UUID) (domain.Wallet, error) {
	var wallet domain.Wallet
	if err := tx.First(&wallet, "id = ?", walletID).Error; err != nil {
		return domain.Wallet{}, err
	}

	return wallet, nil
}

func (wr *walletRepository) UpdateWalletWithTx(tx *gorm.DB, wallet domain.Wallet) error {
	return tx.Model(&domain.Wallet{}).Where("id = ?", wallet.ID).Update("balance", wallet.Balance).Error
}

func (wr *walletRepository) CreateTransaction(transaction domain.Transaction) error {
	return wr.db.Create(&transaction).Error
}

func (wr *walletRepository) GetTransactionsByReferenceID(referenceID uuid.UUID) (domain.Transactions, error) {
	var transactions domain.Transactions
	if err := wr.db.Find(&transactions).Error; err != nil {
		return nil, err
	}

	return transactions, nil
}
