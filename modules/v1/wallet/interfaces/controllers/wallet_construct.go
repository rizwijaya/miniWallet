package controllers

import (
	"github.com/google/uuid"
	"github.com/rizwijaya/miniWallet/modules/common"
	"github.com/rizwijaya/miniWallet/modules/v1/wallet/domain"
)

func constructWalletDisable(wallet domain.Wallet) domain.WalletDisableResponse {
	return domain.WalletDisableResponse{
		ID:         wallet.ID,
		OwnedBy:    wallet.CustomerXID,
		Status:     common.WalletStatusToString[wallet.Status],
		DisabledAt: *wallet.GormModel.UpdatedAt,
		Balance:    wallet.Balance,
	}
}

func constructWalletEnable(wallet domain.Wallet) domain.WalletEnableResponse {
	return domain.WalletEnableResponse{
		ID:        wallet.ID,
		OwnedBy:   wallet.CustomerXID,
		Status:    common.WalletStatusToString[wallet.Status],
		EnabledAt: *wallet.GormModel.UpdatedAt,
		Balance:   wallet.Balance,
	}
}

func constructTransactions(transactions domain.Transactions) (resp []domain.TransactionResponse) {
	for _, transaction := range transactions {
		transactionResponse := domain.TransactionResponse{
			ID:           transaction.ID,
			Status:       common.TransactionStatusToString[transaction.Status],
			TransactedAt: *transaction.CreatedAt,
			Type:         common.TransactionTypeToString[transaction.Type],
			Amount:       transaction.Amount,
			ReferenceID:  transaction.ReferenceID,
		}

		resp = append(resp, transactionResponse)
	}

	return resp
}

func constructDeposit(transaction domain.Transaction, customerXID uuid.UUID) (resp domain.DepositResponse) {
	return domain.DepositResponse{
		ID:          transaction.ID,
		DepositedBy: customerXID,
		Status:      common.TransactionStatusToString[transaction.Status],
		DepositedAt: *transaction.CreatedAt,
		Amount:      transaction.Amount,
		ReferenceID: transaction.ReferenceID,
	}
}

func constructWithdrawal(transaction domain.Transaction, customerXID uuid.UUID) (resp domain.WithdrawalResponse) {
	return domain.WithdrawalResponse{
		ID:           transaction.ID,
		WithdrawalBy: customerXID,
		Status:       common.TransactionStatusToString[transaction.Status],
		WithdrawalAt: *transaction.CreatedAt,
		Amount:       transaction.Amount,
		ReferenceID:  transaction.ReferenceID,
	}
}
