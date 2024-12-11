package controllers

import (
	walletUsecase "github.com/rizwijaya/miniWallet/modules/v1/wallet/usecases"
)

type WalletController struct {
	walletUsecase walletUsecase.WalletUsecase
}

func NewController(walletUsecase walletUsecase.WalletUsecase) *WalletController {
	return &WalletController{
		walletUsecase: walletUsecase,
	}
}
