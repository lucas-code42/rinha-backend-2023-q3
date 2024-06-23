package infra

import (
	"github.com/labstack/echo/v4"
	"github.com/lucas-code42/rinha-backend/infra/handler"
)

func StartHttpServer(e *echo.Echo) {
	e.POST("/pessoas", handler.CreatePerson)
	e.GET("/pessoas/:id", handler.GetPersonById)
	e.GET("/pessoas", handler.SearchPerson)
	e.GET("/contagem-pessoas", handler.CountPeople)

	e.Logger.Fatal(e.Start(":1323"))
}
