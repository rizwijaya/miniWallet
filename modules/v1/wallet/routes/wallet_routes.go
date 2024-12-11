package routes

import (
	"github.com/gofiber/fiber/v2"
	walletCtrl "github.com/rizwijaya/miniWallet/modules/v1/wallet/interfaces/controllers"
)

func Router(ctrl *walletCtrl.WalletController, api fiber.Router) {
	{
		api.Get("/init", ctrl.InitMyAccount)
	}
}
