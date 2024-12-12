package common

const (
	//User sessions
	UserSessionCustomerXID = "customerxid"
	UserSessionWalletID    = "walletid"
	UserSessionExpired     = "expired"
)

const (
	//Wallet status
	WalletStatusActive    = 1
	WalletStatusNonActive = 0

	//Wallet status string
	WalletStatusActiveString    = "enabled"
	WalletStatusNonActiveString = "disabled"
)

const (
	//Transaction status
	TransactionStatusFailed  = 0
	TransactionStatusProcess = 1
	TransactionStatusSuccess = 2

	//Transaction status string
	TransactionStatusFailedString  = "failed"
	TransactionStatusProcessString = "process"
	TransactionStatusSuccessString = "success"
)

const (
	//Transaction type
	TransactionTypeDeposit    = 0
	TransactionTypeWithdrawal = 1

	//Transaction type string
	TransactionTypeDepositString    = "deposit"
	TransactionTypeWithdrawalString = "withdrawal"
)
