package util

import (
	"errors"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/nullrish/task-manager-go/internal/model"
)

func GenerateNewUserToken(userID uuid.UUID, tt string) (string, error) {
	tokenType := model.TokenType(tt)
	expiry := tokenType.GetExpiryTime().Unix()
	var err error
	var token string
	switch tokenType {
	case model.Bearer, model.Refresh:
		claims := jwt.MapClaims{
			"id":  userID.String(),
			"exp": expiry,
		}

		t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		token, err = t.SignedString([]byte(os.Getenv("JWT_KEY")))
		return token, err
	case model.Reset:
		return "refresh", nil
	case model.Verify:
		return "verify", nil
	default:
		return "", errors.New("invalid token type")
	}
}
