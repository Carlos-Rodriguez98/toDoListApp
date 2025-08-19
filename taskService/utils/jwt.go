package utils

import (
	"errors"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func ExtractUserIDFromToken(tokenString, secret string) (uint, error) {
	if tokenString == "" {
		return 0, errors.New("token vacío")
	}
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil || !token.Valid {
		return 0, errors.New("token inválido")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("claims inválidos")
	}

	// el authService guardó "user_id" como número
	switch v := claims["user_id"].(type) {
	case float64:
		return uint(v), nil
	case int:
		return uint(v), nil
	default:
		return 0, errors.New("user_id no presente en token")
	}
}

func AuthMiddleware() gin.HandlerFunc {
	secret := os.Getenv("JWT_SECRET")
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Falta token Bearer"})
			return
		}
		raw := strings.TrimPrefix(authHeader, "Bearer ")
		uid, err := ExtractUserIDFromToken(raw, secret)
		if err != nil || uid == 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			return
		}
		c.Set("user_id", uid)
		c.Next()
	}
}
