package infra

import (
	"fmt"
	"log/slog"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/lucas-code42/rinha-backend/internal/application"
	"github.com/lucas-code42/rinha-backend/internal/application/usecase/create"
	"github.com/lucas-code42/rinha-backend/internal/domain"
)

type HttpController struct {
	respository application.RespositoryImpl
}

func NewController(respository application.RespositoryImpl) *HttpController {
	return &HttpController{
		respository: respository,
	}
}

func (h *HttpController) CreatePerson() func(echo.Context) error {
	return func(c echo.Context) error {
		var payload domain.Pessoa
		if err := c.Bind(&payload); err != nil {
			slog.Error("error payload unprocessable entity", err)
			return c.JSON(http.StatusUnprocessableEntity, map[string]int{"error": http.StatusUnprocessableEntity})
		}

		uuid := uuid.NewString()
		payloadDto := &domain.PessoaDto{
			Id:         uuid,
			Nome:       payload.Nome,
			Apelido:    payload.Apelido,
			Nascimento: payload.Nascimento,
			Stack:      strings.Join(payload.Stack, ";"),
		}

		usecase := create.New(payloadDto, h.respository)
		if err := usecase.Execute(); err != nil {
			return c.JSON(http.StatusUnprocessableEntity, map[string]int{"error": http.StatusUnprocessableEntity})
		}

		c.Response().Header().Add("Location", fmt.Sprintf("/pessoas/%s", uuid))
		return c.JSONBlob(http.StatusCreated, []byte(""))
	}
}

// func CreatePerson(c echo.Context) error {
// 	return c.String(http.StatusInternalServerError, "CreatePerson - NOT IMPLEMENTED!!")
// }

// func GetPersonById(c echo.Context) error {
// 	return c.String(http.StatusInternalServerError, "GetPersonById - NOT IMPLEMENTED!!")
// }

// func SearchPerson(c echo.Context) error {
// 	return c.String(http.StatusInternalServerError, "SearchPerson - NOT IMPLEMENTED!!")
// }

// func CountPeople(c echo.Context) error {
// 	return c.String(http.StatusInternalServerError, "CountPeople - NOT IMPLEMENTED!!")
// }
