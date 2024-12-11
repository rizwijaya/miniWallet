package usecase

import (
	"github.com/google/uuid"
	"github.com/rizwijaya/miniWallet/modules/v1/wallet/domain"
	walletRepository "github.com/rizwijaya/miniWallet/modules/v1/wallet/interfaces/repositories"
)

type WalletUsecase interface {
	InitMyAccount(param domain.InitMyAccountInput) (string, error)
	ChangeStatusWalletByCustomerXID(param domain.ChangeStatusWalletByCustomerXID) (domain.Wallet, error)
	GetWalletByCustomerXID(customerXID uuid.UUID) (domain.Wallet, error)
}

type walletUsecase struct {
	walletRepository walletRepository.WalletRepository
}

func NewUsecase(walletRepository walletRepository.WalletRepository) *walletUsecase {
	return &walletUsecase{
		walletRepository: walletRepository,
	}
}
