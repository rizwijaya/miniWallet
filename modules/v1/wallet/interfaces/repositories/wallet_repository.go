package repository

import (
	"github.com/google/uuid"
	"github.com/rizwijaya/miniWallet/modules/v1/wallet/domain"
)

func (wr *walletRepository) GetWalletByUserID(userID uuid.UUID) (domain.Wallet, error) {
	var wallet domain.Wallet
	if err := wr.db.Model(&domain.Wallet{}).Where("user_id = ?", userID).First(&wallet).Error; err != nil {
		return domain.Wallet{}, err
	}
	return wallet, nil
}

func (wr *walletRepository) CreateWallet(wallet domain.Wallet) error {
	return wr.db.Create(&wallet).Error
}

func (wr *walletRepository) ChangeStatusWalletByUserID(param domain.ChangeStatusWalletByUserID) (domain.Wallet, error) {
	var wallet domain.Wallet
	if err := wr.db.Model(&domain.Wallet{}).Where("user_id = ?", param.UserID).Update("status", param.Status).Scan(&wallet).Error; err != nil {
		return domain.Wallet{}, err
	}

	return wallet, nil
}
