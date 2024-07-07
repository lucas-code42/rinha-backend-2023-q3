package main

import (
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/lucas-code42/rinha-backend/infra"
	"github.com/lucas-code42/rinha-backend/internal/configs"
	"github.com/lucas-code42/rinha-backend/internal/repository"
	"github.com/lucas-code42/rinha-backend/pkg/sql"
)

func main() {
	// TODO: fix this in docker-compose!!!
	time.Sleep(20 * time.Second)
	fmt.Println("START NOW!!")

	configs.Init()
	db := sql.New()
	repo := repository.New(db.SqlClient)
	echo := echo.New()

	infra.New(echo, repo).StartHttpServer()
}
