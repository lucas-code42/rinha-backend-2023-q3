package getpersonbyid

import (
	"log/slog"
	"strings"

	"github.com/lucas-code42/rinha-backend/internal/application"
	"github.com/lucas-code42/rinha-backend/internal/domain"
)

type GetPersonById struct {
	repository application.RespositoryImpl
}

func New(
	repository application.RespositoryImpl,
) *GetPersonById {
	return &GetPersonById{
		repository: repository,
	}
}

func (c *GetPersonById) Execute(personId string) (*domain.Pessoa, error) {
	personDto, err := c.repository.GetPersonById(personId)
	if err != nil {
		slog.Error("error usecase", err)
		return &domain.Pessoa{}, err
	}

	return &domain.Pessoa{
		Id:         personDto.Id,
		Apelido:    personDto.Apelido,
		Nome:       personDto.Nome,
		Nascimento: personDto.Nascimento,
		Stack:      strings.Split(personDto.Stack, ";"),
	}, nil
}
