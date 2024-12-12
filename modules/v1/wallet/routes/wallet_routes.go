package routes

import (
	"github.com/gofiber/fiber/v2"
	middlewareLib "github.com/rizwijaya/miniWallet/infrastructures/middlewares"
	"github.com/rizwijaya/miniWallet/modules/common"
	walletCtrl "github.com/rizwijaya/miniWallet/modules/v1/wallet/interfaces/controllers"
)

func Router(ctrl *walletCtrl.WalletController, api fiber.Router) {
	{
		api.Post("/init", ctrl.InitMyAccount)

		//API using authentication
		api.Use(middlewareLib.Authentication())
		wallet := api.Group("/wallet")
		{
			//API authorization only wallet nonactive can access
			wallet.Post("", middlewareLib.Authorization(common.WalletStatusNonActive), ctrl.EnableMyWallet)

			//API authorization only wallet active can access
			wallet.Use(middlewareLib.Authorization(common.WalletStatusActive))
			wallet.Get("", ctrl.GetWallet)
			wallet.Get("/transactions", ctrl.GetTransactions)
			wallet.Post("/deposits", ctrl.Deposit)
		}
	}
}
