package usecase

import (
	"github.com/google/uuid"
	"github.com/rizwijaya/miniWallet/modules/common"
	"github.com/rizwijaya/miniWallet/modules/v1/wallet/domain"
	errorLib "github.com/rizwijaya/miniWallet/pkg/http_error"
	tokenLib "github.com/rizwijaya/miniWallet/pkg/jwt"
)

func (wuc *walletUsecase) InitMyAccount(param domain.InitMyAccountInput) (string, error) {
	//Check user only have 1 wallet
	wallet, err := wuc.walletRepository.GetWalletByCustomerXID(param.CustomerXID)
	if err != nil && err.Error() != errorLib.ErrRecordNotFound.Error() {
		return "", err
	}

	//user not have wallet, created new wallet
	if err != nil && err.Error() == errorLib.ErrRecordNotFound.Error() {
		wallet = domain.Wallet{
			GormModel: domain.GormModel{
				ID: uuid.New(),
			},
			CustomerXID: param.CustomerXID,
			Balance:     0,
			Status:      common.WalletStatusNonActive,
		}

		err = wuc.walletRepository.CreateWallet(wallet)
		if err != nil {
			return "", err
		}
	}

	// generate token
	accessToken, err := tokenLib.GenerateToken(param.CustomerXID)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

func (wuc *walletUsecase) ChangeStatusWalletByCustomerXID(param domain.ChangeStatusWalletByCustomerXID) (domain.Wallet, error) {
	return wuc.walletRepository.ChangeStatusWalletByCustomerXID(param)
}

func (wuc *walletUsecase) GetWalletByCustomerXID(customerXID uuid.UUID) (domain.Wallet, error) {
	return wuc.walletRepository.GetWalletByCustomerXID(customerXID)
}
