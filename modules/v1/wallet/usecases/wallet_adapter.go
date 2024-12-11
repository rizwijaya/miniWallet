package usecase

import (
	"github.com/rizwijaya/miniWallet/modules/v1/wallet/domain"
	walletRepository "github.com/rizwijaya/miniWallet/modules/v1/wallet/interfaces/repositories"
)

type WalletUsecase interface {
	InitMyAccount(param domain.InitMyAccountInput) (string, error)
}

type walletUsecase struct {
	walletRepository walletRepository.WalletRepository
}

func NewUsecase(walletRepository walletRepository.WalletRepository) *walletUsecase {
	return &walletUsecase{
		walletRepository: walletRepository,
	}
}
