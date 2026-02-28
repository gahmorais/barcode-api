package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/barcode-api/response"
	"github.com/gin-gonic/gin"
)

type userRepoMock struct {
	createFn func(username, password string) error
	loginFn  func(username, password string) error
}

func (u userRepoMock) Create(username, password string) error {
	if u.createFn != nil {
		return u.createFn(username, password)
	}
	return nil
}

func (u userRepoMock) Login(username, password string) error {
	if u.loginFn != nil {
		return u.loginFn(username, password)
	}
	return nil
}

func newContextWithBody(t *testing.T, method string, body any) (*gin.Context, *httptest.ResponseRecorder) {
	t.Helper()

	payload, err := json.Marshal(body)
	if err != nil {
		t.Fatalf("erro ao serializar payload: %v", err)
	}

	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	ctx.Request = httptest.NewRequest(method, "/", bytes.NewReader(payload))
	ctx.Request.Header.Set("Content-Type", "application/json")
	return ctx, recorder
}

func decodeMessage(t *testing.T, recorder *httptest.ResponseRecorder) response.Message {
	t.Helper()

	var message response.Message
	if err := json.Unmarshal(recorder.Body.Bytes(), &message); err != nil {
		t.Fatalf("erro ao decodificar resposta: %v", err)
	}

	return message
}

func TestCreateUser(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("deve criar usuário com sucesso", func(t *testing.T) {
		called := false
		controller := NewUserController(userRepoMock{createFn: func(username, password string) error {
			called = true
			if username != "maria" || password != "123456789" {
				t.Fatalf("payload enviado ao repository está incorreto")
			}
			return nil
		}})

		ctx, recorder := newContextWithBody(t, http.MethodPost, map[string]string{
			"username": "maria",
			"password": "123456789",
		})

		controller.CreateUser(ctx)

		if !called {
			t.Fatalf("esperava chamada ao repository")
		}
		if recorder.Code != http.StatusOK {
			t.Fatalf("status inesperado: %d", recorder.Code)
		}

		message := decodeMessage(t, recorder)
		if message.Text != "Usuário criado com sucesso" {
			t.Fatalf("mensagem inesperada: %s", message.Text)
		}
	})

	t.Run("deve retornar bad request para payload inválido", func(t *testing.T) {
		controller := NewUserController(userRepoMock{})
		recorder := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(recorder)
		ctx.Request = httptest.NewRequest(http.MethodPost, "/", bytes.NewBufferString("{invalid"))

		controller.CreateUser(ctx)

		if recorder.Code != http.StatusInternalServerError {
			t.Fatalf("status inesperado: %d", recorder.Code)
		}
	})

	t.Run("deve rejeitar usuário ou senha vazios", func(t *testing.T) {
		called := false
		controller := NewUserController(userRepoMock{createFn: func(username, password string) error {
			called = true
			return nil
		}})

		ctx, recorder := newContextWithBody(t, http.MethodPost, map[string]string{
			"username": "",
			"password": "123456789",
		})

		controller.CreateUser(ctx)

		if called {
			t.Fatalf("repository não deveria ser chamado")
		}
		if recorder.Code != http.StatusBadRequest {
			t.Fatalf("status inesperado: %d", recorder.Code)
		}
	})

	t.Run("deve rejeitar senha curta", func(t *testing.T) {
		called := false
		controller := NewUserController(userRepoMock{createFn: func(username, password string) error {
			called = true
			return nil
		}})

		ctx, recorder := newContextWithBody(t, http.MethodPost, map[string]string{
			"username": "maria",
			"password": "1234567",
		})

		controller.CreateUser(ctx)

		if called {
			t.Fatalf("repository não deveria ser chamado")
		}
		if recorder.Code != http.StatusBadRequest {
			t.Fatalf("status inesperado: %d", recorder.Code)
		}
	})

	t.Run("deve retornar erro interno quando repository falha", func(t *testing.T) {
		controller := NewUserController(userRepoMock{createFn: func(username, password string) error {
			return errors.New("erro de persistência")
		}})

		ctx, recorder := newContextWithBody(t, http.MethodPost, map[string]string{
			"username": "maria",
			"password": "123456789",
		})

		controller.CreateUser(ctx)

		if recorder.Code != http.StatusInternalServerError {
			t.Fatalf("status inesperado: %d", recorder.Code)
		}
	})
}

func TestLogin(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("deve logar com sucesso", func(t *testing.T) {
		called := false
		controller := NewUserController(userRepoMock{loginFn: func(username, password string) error {
			called = true
			if username != "maria" || password != "123456789" {
				t.Fatalf("payload enviado ao repository está incorreto")
			}
			return nil
		}})

		ctx, recorder := newContextWithBody(t, http.MethodPost, map[string]string{
			"UserName": "maria",
			"Password": "123456789",
		})

		controller.Login(ctx)

		if !called {
			t.Fatalf("esperava chamada ao repository")
		}
		if recorder.Code != http.StatusOK {
			t.Fatalf("status inesperado: %d", recorder.Code)
		}
	})

	t.Run("deve retornar unauthorized quando login falha", func(t *testing.T) {
		controller := NewUserController(userRepoMock{loginFn: func(username, password string) error {
			return errors.New("usuário ou senha incorretos")
		}})

		ctx, recorder := newContextWithBody(t, http.MethodPost, map[string]string{
			"UserName": "maria",
			"Password": "senhaerrada",
		})

		controller.Login(ctx)

		if recorder.Code != http.StatusUnauthorized {
			t.Fatalf("status inesperado: %d", recorder.Code)
		}
	})

	t.Run("deve retornar erro para json inválido", func(t *testing.T) {
		controller := NewUserController(userRepoMock{})
		recorder := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(recorder)
		ctx.Request = httptest.NewRequest(http.MethodPost, "/", bytes.NewBufferString("{invalid"))

		controller.Login(ctx)

		if recorder.Code != http.StatusInternalServerError {
			t.Fatalf("status inesperado: %d", recorder.Code)
		}
	})
}
