package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
	"github.com/rizwijaya/miniWallet/modules/common"
	"github.com/rizwijaya/miniWallet/modules/v1/wallet/domain"
	apiResponse "github.com/rizwijaya/miniWallet/pkg/api_response"
)

func (wc *WalletController) InitMyAccount(c *fiber.Ctx) error {
	var (
		req  domain.InitMyAccountInput
		resp apiResponse.Response
	)

	defer func() {
		log.Debugf("[INCOMING REQUEST INIT MY ACCOUNT][%s][IP: %s][REQ: %s][RESP: %s]", c.Method(), c.IP(), common.MustMarshal(req), common.MustMarshal(resp))
	}()

	if err := c.BodyParser(&req); err != nil {
		log.Errorf("[ERROR][InitMyAccount][BodyParser][%s]", err.Error())
		resp = apiResponse.CustomResponse(fmt.Sprintf("Failed to parse request: %s", err.Error()), apiResponse.HttpStatusFailed)
		return c.Status(fiber.StatusBadRequest).JSON(resp)
	}

	if req.CustomerXID == uuid.Nil {
		log.Errorf("[ERROR][InitMyAccount][CustomerXIDEmpty][Missing data for required field.]")
		resp = apiResponse.CustomResponse(map[string][]string{
			"customer_xid": {"Missing data for required field."},
		}, apiResponse.HttpStatusFailed)

		return c.Status(fiber.StatusBadRequest).JSON(resp)
	}

	token, err := wc.walletUsecase.InitMyAccount(req)
	if err != nil {
		log.Errorf("[ERROR][InitMyAccount][uc:InitMyAccount][%s]", err.Error())
		resp = apiResponse.CustomResponse(err.Error(), apiResponse.HttpStatusFailed)
		return c.Status(fiber.StatusBadRequest).JSON(resp)
	}

	resp = apiResponse.CustomResponse(map[string]string{
		"token": token,
	}, apiResponse.HttpStatusSuccess)

	return c.Status(fiber.StatusOK).JSON(resp)
}

func (wc *WalletController) EnableMyWallet(c *fiber.Ctx) error {
	var (
		resp        apiResponse.Response
		customerXID = c.Locals(common.UserSessionCustomerXID).(uuid.UUID)
	)

	defer func() {
		log.Debugf("[INCOMING REQUEST ENABLE MY WALLET][%s][IP: %s][CUSTOMERXID: %s][RESP: %s]", c.Method(), c.IP(), customerXID, common.MustMarshal(resp))
	}()

	wallet, err := wc.walletUsecase.ChangeStatusWalletByCustomerXID(domain.ChangeStatusWalletByCustomerXID{
		CustomerXID: customerXID,
		Status:      common.WalletStatusActive,
	})
	if err != nil {
		log.Errorf("[ERROR][EnableMyWallet][uc:ChangeStatusWalletByCustomerXID][%s]", err.Error())
		resp = apiResponse.CustomResponse(err.Error(), apiResponse.HttpStatusFailed)
		return c.Status(fiber.StatusBadRequest).JSON(resp)
	}

	resp = apiResponse.CustomResponse(map[string]interface{}{
		"wallet": constructWallet(wallet),
	}, apiResponse.HttpStatusSuccess)

	return c.Status(fiber.StatusCreated).JSON(resp)
}

func (wc *WalletController) GetWallet(c *fiber.Ctx) error {
	var (
		resp        apiResponse.Response
		customerXID = c.Locals(common.UserSessionCustomerXID).(uuid.UUID)
	)

	defer func() {
		log.Debugf("[INCOMING REQUEST GET WALLET][%s][IP: %s][CUSTOMERXID: %s][RESP: %s]", c.Method(), c.IP(), customerXID, common.MustMarshal(resp))
	}()

	wallet, err := wc.walletUsecase.GetWalletByCustomerXID(customerXID)
	if err != nil {
		log.Errorf("[ERROR][GetWallet][uc:GetWalletByCustomerXID][%s]", err.Error())
		resp = apiResponse.CustomResponse(err.Error(), apiResponse.HttpStatusFailed)
		return c.Status(fiber.StatusBadRequest).JSON(resp)
	}

	resp = apiResponse.CustomResponse(map[string]interface{}{
		"wallet": constructWallet(wallet),
	}, apiResponse.HttpStatusSuccess)

	return c.Status(fiber.StatusOK).JSON(resp)
}
