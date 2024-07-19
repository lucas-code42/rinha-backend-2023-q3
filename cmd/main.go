package main

import (
	"github.com/labstack/echo/v4"
	"github.com/lucas-code42/rinha-backend/infra"
	"github.com/lucas-code42/rinha-backend/internal/configs"
	"github.com/lucas-code42/rinha-backend/internal/repository"
	"github.com/lucas-code42/rinha-backend/pkg/sql"
)

func main() {
	configs.Init()
	db := sql.New()
	repo := repository.New(db.SqlClient)
	echo := echo.New()

	srv := infra.New(echo, repo)
	srv.SetupRouters()
	srv.StartServer()
}
