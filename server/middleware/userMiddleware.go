package middleware

import (
	"net/http"

	"ashwin.com/go-crypto-simulator/helper"
	"github.com/gin-gonic/gin"
)

func AuthenticateMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, tokenErr := c.Cookie("token")
		if tokenErr != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Cookie error"})
			c.Abort()
			return
		}

		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token error"})
			c.Abort()
			return
		}

		claims, errMsg := helper.ValidateToken(token)
		if errMsg != "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": errMsg})
			c.Abort()
			return
		}

		c.Set("name", claims.Name)
		c.Set("id", claims.UID)

		c.Next()
	}
}
