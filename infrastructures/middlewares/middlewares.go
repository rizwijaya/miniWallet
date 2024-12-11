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
		if err := database.Model(&domain.Wallet{}).Where("user_id = ?", c.Locals(common.UserSessionUserID).(uuid.UUID)).First(&wallet).Error; err != nil {
			log.Infof("[Middleware][PATH: %s][%s][IP: %s][ERROR: %s]", c.Path(), c.Method(), c.IP(), err)
			return c.Status(fiber.StatusUnauthorized).JSON(apiResponse.CustomResponse("Unauthorized", apiResponse.HttpStatusFailed))
		}

		if wallet.Status != walletStatus {
			log.Infof("[Middleware][PATH: %s][%s][IP: %s][ERROR: %s]", c.Path(), c.Method(), c.IP(), fmt.Sprintf("Wallet is Already %s", common.WalletStatusToString[wallet.Status]))
			return c.Status(fiber.StatusBadRequest).JSON(apiResponse.CustomResponse(fmt.Sprintf("Already %s", common.WalletStatusToString[wallet.Status]), apiResponse.HttpStatusFailed))
		}

		return c.Next()
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
		userID := uuid.MustParse(token.Claims.(jwt.MapClaims)[common.UserSessionUserID].(string))

		// Save userSessions to context
		c.Locals(common.UserSessionUserID, userID)
		return c.Next()
	}
}
