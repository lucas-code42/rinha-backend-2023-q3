package infra

import (
	"github.com/labstack/echo/v4"
	"github.com/lucas-code42/rinha-backend/internal/application"
)

type HttpServer struct {
	echoEngine *echo.Echo
	repository application.RespositoryImpl
}

func New(
	echoEngine *echo.Echo,
	repository application.RespositoryImpl,
) *HttpServer {
	return &HttpServer{
		echoEngine: echoEngine,
		repository: repository,
	}
}

func (h *HttpServer) StartHttpServer() {
	httpController := NewController(h.repository)

	h.echoEngine.POST("/pessoas", httpController.CreatePerson())
	// h.echoEngine.GET("/pessoas/:id", httpController.GetPersonById())
	// h.echoEngine.GET("/pessoas", httpController.SearchPerson)
	// h.echoEngine.GET("/contagem-pessoas", httpController.CountPeople)

	h.echoEngine.Logger.Fatal(h.echoEngine.Start(":8080"))
}
