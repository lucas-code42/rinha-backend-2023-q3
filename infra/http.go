package infra

import (
	"github.com/labstack/echo/v4"
	"github.com/lucas-code42/rinha-backend/internal/application"
)

type HttpServer struct {
	EchoEngine *echo.Echo
	repository application.RespositoryImpl
}

func New(
	echoEngine *echo.Echo,
	repository application.RespositoryImpl,
) *HttpServer {
	return &HttpServer{
		EchoEngine: echoEngine,
		repository: repository,
	}
}

func (h *HttpServer) SetupRouters() *echo.Echo {
	httpController := NewController(h.repository)

	h.EchoEngine.GET("/live", func(c echo.Context) error {
		return c.JSON(200, "OK")
	})

	h.EchoEngine.POST("/pessoas", httpController.CreatePerson())
	h.EchoEngine.GET("/pessoas/:id", httpController.GetPersonById())
	h.EchoEngine.GET("/pessoas", httpController.SearchPerson())
	h.EchoEngine.GET("/contagem-pessoas", httpController.CountPeople())

	return h.EchoEngine
}

func (h *HttpServer) StartServer() {
	h.EchoEngine.Start(":80")
}
