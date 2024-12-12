package common

var (
	//Convert Wallet Status
	WalletStatusToString = map[int]string{
		WalletStatusActive:    WalletStatusActiveString,
		WalletStatusNonActive: WalletStatusNonActiveString,
	}

	//Convert Transaction Status
	TransactionStatusToString = map[int]string{
		TransactionStatusFailed:  TransactionStatusFailedString,
		TransactionStatusProcess: TransactionStatusProcessString,
		TransactionStatusSuccess: TransactionStatusSuccessString,
	}

	//Convert Transaction Type
	TransactionTypeToString = map[int]string{
		TransactionTypeDeposit:    TransactionTypeDepositString,
		TransactionTypeWithdrawal: TransactionTypeWithdrawalString,
	}
)
