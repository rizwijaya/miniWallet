package token

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/rizwijaya/miniWallet/infrastructures/config"
	"github.com/rizwijaya/miniWallet/modules/common"
	timeLib "github.com/rizwijaya/miniWallet/pkg/time"
)

func GenerateToken(userID uuid.UUID) (string, error) {
	config, err := config.New()
	if err != nil {
		return "", err
	}

	claim := jwt.MapClaims{}
	claim[common.UserSessionUserID] = userID
	claim[common.UserSessionExpired] = timeLib.TimeNow().Add(time.Minute * 30).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	secretKey := []byte(config.App.Secret_key)

	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
