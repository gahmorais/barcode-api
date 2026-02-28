package auth

import (
	"net/http"
	"strings"

	"github.com/barcode-api/response"
	"github.com/gin-gonic/gin"
)

func Authentication(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, response.Message{Text: "authorization header ausente"})
			c.Abort()
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, response.Message{Text: "token bearer obrigatório"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := VerifyToken(tokenString, secret)
		if err != nil {
			c.JSON(http.StatusUnauthorized, response.Message{Text: "token inválido"})
			c.Abort()
			return
		}

		c.Set("username", claims["username"])
		c.Next()
	}
}
