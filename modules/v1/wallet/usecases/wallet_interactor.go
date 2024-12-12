package usecase

import (
	"errors"

	"github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
	"github.com/rizwijaya/miniWallet/modules/common"
	"github.com/rizwijaya/miniWallet/modules/v1/wallet/domain"
	errorLib "github.com/rizwijaya/miniWallet/pkg/http_error"
	tokenLib "github.com/rizwijaya/miniWallet/pkg/jwt"
	timeLib "github.com/rizwijaya/miniWallet/pkg/time"
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
	accessToken, err := tokenLib.GenerateToken(param.CustomerXID, wallet.ID)
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

func (wuc *walletUsecase) GetTransactionsByCustomerXID(customerXID uuid.UUID) (domain.Transactions, error) {
	return wuc.walletRepository.GetTransactionsByCustomerXID(customerXID)
}

func (wuc *walletUsecase) Deposit(param domain.Deposit) (domain.Transaction, error) {
	//Validate referenceID is unique
	transactionReference, err := wuc.walletRepository.GetTransactionsByReferenceID(param.ReferenceID)
	if err != nil {
		return domain.Transaction{}, err
	}
	if len(transactionReference) > 0 {
		return domain.Transaction{}, errors.New("reference_id used")
	}

	//Get DB Transaction
	tx := wuc.walletRepository.GetDBTx()
	if err = tx.Error; err != nil {
		return domain.Transaction{}, err
	}

	transaction := domain.Transaction{
		GormModel: domain.GormModel{
			ID: uuid.New(),
		},
		WalletID:    param.WalletID,
		Type:        common.TransactionTypeDeposit,
		Amount:      param.Amount,
		ReferenceID: param.ReferenceID,
		Status:      common.TransactionStatusProcess,
	}

	//Rollback when failed to create transaction
	defer func() {
		if r := recover(); r != nil || err != nil {
			tx.Rollback()
			//Update Transaction status to failed
			transaction.Status = common.TransactionStatusFailed
			log.Errorf("[ERROR][uc:Deposit][Rollback][%v %v]", r)
		}

		//Add Transaction
		err = wuc.walletRepository.CreateTransaction(transaction)
		if err != nil {
			tx.Rollback()
			log.Errorf("[ERROR][uc:Deposit][Rollback][CreateTransaction]")
			return
		}

		// Commit the transaction
		if err := tx.Commit().Error; err != nil {
			log.Errorf("[ERROR][uc:Deposit][Rollback][Commit]")
			return
		}

		return
	}()

	//Get Balance from wallet by WalletID
	wallet, err := wuc.walletRepository.GetWalletByIDWithTx(tx, param.WalletID)
	if err != nil {
		transaction.Status = common.TransactionStatusFailed
		return domain.Transaction{}, err
	}

	//Update balance at wallet
	err = wuc.walletRepository.UpdateWalletWithTx(tx, domain.Wallet{
		GormModel: domain.GormModel{
			ID: param.WalletID,
		},
		Balance: wallet.Balance + param.Amount,
	})
	if err != nil {
		transaction.Status = common.TransactionStatusFailed
		return domain.Transaction{}, err
	}

	time := timeLib.TimeNow()
	transaction.CreatedAt = &time
	transaction.UpdatedAt = &time
	transaction.Status = common.TransactionStatusSuccess

	return transaction, nil
}
