package searchperson

import (
	"log/slog"

	"github.com/lucas-code42/rinha-backend/internal/application"
	"github.com/lucas-code42/rinha-backend/internal/domain"
)

type SearchPerson struct {
	repository application.RespositoryImpl
}

func New(
	repository application.RespositoryImpl,
) *SearchPerson {
	return &SearchPerson{
		repository: repository,
	}
}

func (c *SearchPerson) Execute(searchTerm string) ([]*domain.PessoaDto, error) {
	personDto, err := c.repository.SearchPerson(searchTerm)
	if err != nil {
		slog.Error("error usecase", err)
		return []*domain.PessoaDto{}, err
	}

	return personDto, nil
}
