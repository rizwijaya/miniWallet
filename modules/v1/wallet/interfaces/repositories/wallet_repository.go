package repository

import (
	"github.com/google/uuid"
	"github.com/rizwijaya/miniWallet/modules/v1/wallet/domain"
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
