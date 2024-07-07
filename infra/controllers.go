package infra

import (
	"fmt"
	"log/slog"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/lucas-code42/rinha-backend/internal/application"
	"github.com/lucas-code42/rinha-backend/internal/application/usecase/countpeople"
	"github.com/lucas-code42/rinha-backend/internal/application/usecase/createperson"
	"github.com/lucas-code42/rinha-backend/internal/application/usecase/getpersonbyid"
	"github.com/lucas-code42/rinha-backend/internal/application/usecase/searchperson"
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

// validar campos de entrada...
func (h *HttpController) CreatePerson() func(echo.Context) error {
	return func(c echo.Context) error {
		var payload domain.Pessoa
		if err := c.Bind(&payload); err != nil {
			slog.Error("error payload unprocessable entity", err.Error(), err)
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

		usecase := createperson.New(payloadDto, h.respository)
		if err := usecase.Execute(); err != nil {
			slog.Error("error usecase CreatePerson", err.Error(), err)
			return c.JSON(http.StatusUnprocessableEntity, map[string]int{"error": http.StatusUnprocessableEntity})
		}

		c.Response().Header().Add("Location", fmt.Sprintf("/pessoas/%s", uuid))
		return c.JSONBlob(http.StatusCreated, nil)
	}
}

func (h *HttpController) GetPersonById() func(echo.Context) error {
	return func(c echo.Context) error {
		personId := c.Param("id")
		if personId == "" {
			slog.Error("error path param is empty")
			return c.JSON(http.StatusBadRequest, map[string]int{"error": http.StatusBadRequest})
		}

		getUc := getpersonbyid.New(h.respository)
		person, err := getUc.Execute(personId)
		if err != nil {
			slog.Error("error usecase GetPersonById", err.Error(), err)
			return c.JSON(http.StatusInternalServerError, map[string]int{"error": http.StatusInternalServerError})
		}

		return c.JSON(http.StatusOK, person)
	}

}

func (h *HttpController) SearchPerson() func(echo.Context) error {
	return func(c echo.Context) error {
		searchTerm := c.QueryParam("t")
		if searchTerm == "" {
			slog.Error("error query param is empty")
			return c.JSON(http.StatusBadRequest, map[string]int{"error": http.StatusBadRequest})
		}

		searchPersonUc := searchperson.New(h.respository)
		pagination, err := searchPersonUc.Execute(searchTerm)
		if err != nil {
			slog.Error("error usecase SearchPerson", err.Error(), err)
			return c.JSON(http.StatusInternalServerError, map[string]int{"error": http.StatusInternalServerError})
		}

		return c.JSON(http.StatusOK, pagination)
	}
}

func (h *HttpController) CountPeople() func(echo.Context) error {
	return func(c echo.Context) error {
		countUc := countpeople.New(h.respository)
		total, err := countUc.Execute()
		if err != nil {
			slog.Error("error usecase CountPeople")
			return c.JSONBlob(http.StatusInternalServerError, nil)
		}

		return c.JSON(http.StatusOK, map[string]int{"totalRecords": total})
	}
}
