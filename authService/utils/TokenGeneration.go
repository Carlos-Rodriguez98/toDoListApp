package utils

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateJWT(userID uint, userName string) (string, error) {
	claims := jwt.MapClaims{}
	claims["user_id"] = userID
	claims["user_name"] = userName
	claims["expiration"] = time.Now().Add(time.Hour * 24).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}
