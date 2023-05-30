package utils

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

const BLANK = ""

func AuthMiddleware(c *gin.Context) {

	tokenStr := strings.TrimPrefix(c.GetHeader("Authorization"), "Bearer ")

	if tokenStr != BLANK {

		token, err := VerifyToken(tokenStr)

		if err == nil && token.Valid {
			claims, ok := token.Claims.(jwt.MapClaims)
			if ok {
				c.Set("uid", claims["userID"])
				c.Next()
				return
			}
		}
	}

	c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication Required!"})
	c.Abort()
}
