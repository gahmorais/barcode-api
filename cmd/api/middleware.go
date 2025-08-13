package main

import (
	"net/http"
	"strings"

	"github.com/barcode-api/repository"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func (app *application) AuthMiddleware(userRepository repository.UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, Error{Message: "Authorization header is required"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			c.JSON(http.StatusUnauthorized, Error{Message: "Bearer token is required"})
			c.Abort()
			return
		}
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(app.jwtSecret), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, Error{Message: "invalid token"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, Error{Message: "Bearer token is required"})
			c.Abort()
			return
		}

		username := claims["username"].(string)
		user, err := userRepository.GetByUsername(username)
		if err != nil {
			c.JSON(http.StatusUnauthorized, Error{Message: "Unauthorized access"})
			c.Abort()
			return
		}

		c.Set("user", user)
		c.Next()
	}
}
