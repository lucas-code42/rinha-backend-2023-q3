package searchperson

import (
	"log/slog"
	"strings"

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

func (c *SearchPerson) Execute(searchTerm string) ([]*domain.Pessoa, error) {
	personDto, err := c.repository.SearchPerson(searchTerm)
	if err != nil {
		slog.Error("error usecase", err)
		return []*domain.Pessoa{}, err
	}

	if len(personDto) == 0 {
		slog.Error("error usecase", err)
		return []*domain.Pessoa{}, nil
	}

	return c.ParseDto(personDto), nil
}

func (c *SearchPerson) ParseDto(dto []*domain.PessoaDto) []*domain.Pessoa {
	var people []*domain.Pessoa
	for _, v := range dto {
		person := &domain.Pessoa{
			Id:         v.Id,
			Apelido:    v.Apelido,
			Nome:       v.Nome,
			Nascimento: v.Nascimento,
			Stack:      strings.Split(v.Stack, ";"),
		}
		people = append(people, person)
	}

	return people
}
