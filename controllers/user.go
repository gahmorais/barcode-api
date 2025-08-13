package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/barcode-api/models"
	"github.com/barcode-api/repository"
	"github.com/barcode-api/response"
	"github.com/gin-gonic/gin"
)

type User struct {
	repository repository.UserRepository
}

func NewUserController(repo repository.UserRepository) *User {
	return &User{
		repository: repo,
	}
}

func (u *User) CreateUser(c *gin.Context) {
	body := c.Request.Body
	bodyContent, err := io.ReadAll(body)

	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Message{
			Text: "erro ao ler conteúdo do body",
		})
		return
	}

	var user models.UserResponse
	if err := json.Unmarshal(bodyContent, &user); err != nil {
		c.JSON(http.StatusInternalServerError, response.Message{
			Text: "erro json",
		})
		return
	}
	if user.Password == "" || user.Username == "" {
		c.JSON(http.StatusBadRequest, response.Message{
			Text: "usuário e/ou senha não podem ser vazios",
		})
	}
	if len(user.Password) <= 8 {
		c.JSON(http.StatusBadRequest, response.Message{
			Text: "a senha precisa ter mais de 8 caracteres",
		})
		return
	}

	if err := u.repository.Create(user.Username, user.Password); err != nil {
		fmt.Printf("ocorreu um erro : %s", err.Error())
		c.JSON(http.StatusInternalServerError, response.Message{
			Text: "erro ao criar usuário",
		})
		return
	}

	c.JSON(http.StatusOK, response.Message{
		Text: "Usuário criado com sucesso",
	})
}

func (u *User) Login(c *gin.Context) {
	body := c.Request.Body
	bodyContent, err := io.ReadAll(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Message{
			Text: "erro ao ler o body da requisição",
		})
		return
	}
	var user models.User
	if err := json.Unmarshal(bodyContent, &user); err != nil {
		c.JSON(http.StatusInternalServerError, response.Message{
			Text: "erro ao decodificar json",
		})
		return
	}

	if err := u.repository.Login(user.UserName, user.Password); err != nil {
		c.JSON(http.StatusUnauthorized, response.Message{
			Text: err.Error(),
		})
		return
	}

	//gerar token jwt e devolver ao usuário
	c.JSON(http.StatusOK, response.Message{
		Text: "Login Ok",
	})
}
