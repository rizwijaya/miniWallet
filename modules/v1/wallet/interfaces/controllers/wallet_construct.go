package controllers

import (
	"github.com/rizwijaya/miniWallet/modules/common"
	"github.com/rizwijaya/miniWallet/modules/v1/wallet/domain"
)

func constructWallet(wallet domain.Wallet) domain.WalletResponse {
	return domain.WalletResponse{
		ID:        wallet.ID,
		OwnedBy:   wallet.CustomerXID,
		Status:    common.WalletStatusToString[wallet.Status],
		EnabledAt: *wallet.GormModel.UpdatedAt,
		Balance:   wallet.Balance,
	}
}
