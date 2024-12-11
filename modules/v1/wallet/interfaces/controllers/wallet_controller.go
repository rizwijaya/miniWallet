package controllers

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
	"github.com/rizwijaya/miniWallet/modules/common"
	"github.com/rizwijaya/miniWallet/modules/v1/wallet/domain"
	api "github.com/rizwijaya/miniWallet/pkg/api_response"
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
		resp = api.CustomResponse(fmt.Sprintf("Failed to parse request: %s", err.Error()), apiResponse.HttpStatusFailed)
		return c.Status(http.StatusBadRequest).JSON(resp)
	}

	if req.UserID == uuid.Nil {
		log.Errorf("[ERROR][InitMyAccount][UserIDEmpty][Missing data for required field.]")
		resp = api.CustomResponse(map[string][]string{
			"customer_xid": {"Missing data for required field."},
		}, apiResponse.HttpStatusFailed)

		return c.Status(http.StatusBadRequest).JSON(resp)
	}

	token, err := wc.walletUsecase.InitMyAccount(req)
	if err != nil {
		log.Errorf("[ERROR][InitMyAccount][uc:InitMyAccount][%s]", err.Error())
		resp = api.CustomResponse(err.Error(), apiResponse.HttpStatusFailed)
		return c.Status(http.StatusBadRequest).JSON(resp)
	}

	resp = api.CustomResponse(map[string]string{
		"token": token,
	}, apiResponse.HttpStatusSuccess)

	return c.Status(http.StatusOK).JSON(resp)
}
