package main

import (
	"github.com/labstack/echo/v4"
	"github.com/lucas-code42/rinha-backend/infra"
	"github.com/lucas-code42/rinha-backend/internal/configs"
	"github.com/lucas-code42/rinha-backend/pkg/database"
)

type Pessoa struct {
	Id         string   `json:"id,omitempty"`
	Apelido    string   `json:"apelido,omitempty"`
	Nome       string   `json:"nome,omitempty"`
	Nascimento string   `json:"nascimento,omitempty"`
	Stack      []string `json:"stack,omitempty"`
}

func main() {
	configs := configs.New()
	echoEngine := echo.New()
	db := database.New()

	infra.StartHttpServer(echoEngine)
}
