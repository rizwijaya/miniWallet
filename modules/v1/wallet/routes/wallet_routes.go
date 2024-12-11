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

		//API Using Authentication
		api.Use(middlewareLib.Authentication())
		api.Post("/wallet", middlewareLib.Authorization(common.WalletStatusNonActive), ctrl.EnableMyWallet)

	}
}
