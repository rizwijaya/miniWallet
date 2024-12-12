package middlewares

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
	"github.com/rizwijaya/miniWallet/modules/common"
	"github.com/rizwijaya/miniWallet/modules/v1/wallet/domain"
	apiResponse "github.com/rizwijaya/miniWallet/pkg/api_response"
	tokenLib "github.com/rizwijaya/miniWallet/pkg/jwt"
	"gorm.io/gorm"
)

var database *gorm.DB

func NewMiddleware(db *gorm.DB) {
	database = db
}

func Authorization(walletStatus int) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var wallet domain.Wallet
		if err := database.Model(&domain.Wallet{}).Where("customer_xid = ?", c.Locals(common.UserSessionCustomerXID).(uuid.UUID)).First(&wallet).Error; err != nil {
			log.Infof("[Middleware][PATH: %s][%s][IP: %s][ERROR: %s]", c.Path(), c.Method(), c.IP(), err)
			return c.Status(fiber.StatusUnauthorized).JSON(apiResponse.CustomResponse("Unauthorized", apiResponse.HttpStatusFailed))
		}

		if wallet.Status == walletStatus {
			return c.Next()
		}

		//Custom Response for specific PATH
		logDesc := fmt.Sprintf("Wallet %s", common.WalletStatusToString[wallet.Status])
		codeResponse := fiber.StatusNotFound
		if c.Path() == "/api/v1/wallet" && (c.Method() == "POST" || c.Method() == "PATCH") {
			logDesc = fmt.Sprintf("Already %s", common.WalletStatusToString[wallet.Status])
			codeResponse = fiber.StatusBadRequest
		}

		log.Infof("[Middleware][PATH: %s][%s][IP: %s][ERROR: %s]", c.Path(), c.Method(), c.IP(), logDesc)
		return c.Status(codeResponse).JSON(apiResponse.CustomResponse(logDesc, apiResponse.HttpStatusFailed))
	}
}

func Authentication() fiber.Handler {
	return func(c *fiber.Ctx) error {
		//Extract and validate token
		// Get token from header
		tokenString := c.Get("Authorization")
		if tokenString == "" {
			log.Infof("[Middleware][PATH: %s][%s][IP: %s][ERROR: %s]", c.Path(), c.Method(), c.IP(), "No token provided")
			return c.Status(fiber.StatusUnauthorized).JSON(apiResponse.CustomResponse("Unauthorized", apiResponse.HttpStatusFailed))
		}

		// Remove "Token " prefix
		if len(tokenString) > 6 && tokenString[:6] == "Token " {
			tokenString = tokenString[6:]
		}

		// Validate token
		token, err := tokenLib.ValidateToken(tokenString)
		if err != nil || !token.Valid {
			log.Infof("[Middleware][PATH: %s][%s][IP: %s][ERROR: %s]", c.Path(), c.Method(), c.IP(), "Invalid token")
			return c.Status(fiber.StatusUnauthorized).JSON(apiResponse.CustomResponse("Unauthorized", apiResponse.HttpStatusFailed))
		}
		customerXID := uuid.MustParse(token.Claims.(jwt.MapClaims)[common.UserSessionCustomerXID].(string))
		walletID := uuid.MustParse(token.Claims.(jwt.MapClaims)[common.UserSessionWalletID].(string))

		// Save userSessions to context
		c.Locals(common.UserSessionCustomerXID, customerXID)
		c.Locals(common.UserSessionWalletID, walletID)

		return c.Next()
	}
}
