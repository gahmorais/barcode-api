package auth

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/barcode-api/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Authentication(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	body := c.Request.Body
	bodyBytes, err := io.ReadAll(body)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}
	json.Unmarshal(bodyBytes, &user)

	userDefault := models.User{
		UserName: "Gabriel",
		Password: "123456",
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userDefault.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "username or password invalid",
		})
	}

	newToken, err := GenerateTokenJwt()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Erro ao gerar o token",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"token": newToken,
	})
}
