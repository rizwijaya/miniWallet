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

		//API authorization only wallet nonactive can access
		api.Post("/wallet", middlewareLib.Authorization(common.WalletStatusNonActive), ctrl.EnableMyWallet)

		//API authorization only wallet active can access
		api.Use(middlewareLib.Authorization(common.WalletStatusActive))
		api.Get("/wallet", ctrl.GetWallet)
	}
}
