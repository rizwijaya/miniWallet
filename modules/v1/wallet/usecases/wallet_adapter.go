package usecase

import (
	walletRepository "github.com/rizwijaya/miniWallet/modules/v1/wallet/interfaces/repositories"
)

type WalletUsecase interface {
}

type walletUsecase struct {
	walletRepository walletRepository.WalletRepository
}

func NewUsecase(walletRepository walletRepository.WalletRepository) *walletUsecase {
	return &walletUsecase{
		walletRepository: walletRepository,
	}
}
