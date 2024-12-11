package token

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/rizwijaya/miniWallet/infrastructures/config"
	"github.com/rizwijaya/miniWallet/modules/common"
	timeLib "github.com/rizwijaya/miniWallet/pkg/time"
)

func ValidateToken(encodedToken string) (*jwt.Token, error) {
	conf, err := config.New()
	if err != nil {
		return nil, err
	}
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid token")
		}

		return []byte(conf.App.Secret_key), nil
	})

	if err != nil {
		return token, err
	}

	exp := token.Claims.(jwt.MapClaims)[common.UserSessionExpired].(float64)
	if time.Unix(int64(exp), 0).Before(time.Now()) {
		return token, errors.New("token is expired")
	}

	return token, nil
}

func GenerateToken(customerXID uuid.UUID) (string, error) {
	config, err := config.New()
	if err != nil {
		return "", err
	}

	claim := jwt.MapClaims{}
	claim[common.UserSessionCustomerXID] = customerXID
	claim[common.UserSessionExpired] = timeLib.TimeNow().Add(time.Minute * 30).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	secretKey := []byte(config.App.Secret_key)

	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
