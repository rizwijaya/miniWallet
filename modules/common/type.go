package common

var (
	//Convert Wallet Status
	WalletStatusToString = map[int]string{
		WalletStatusActive:    "enabled",
		WalletStatusNonActive: "disabled",
	}
)
