package application

import "github.com/lucas-code42/rinha-backend/internal/domain"

type RespositoryImpl interface {
	CreatePerson(person *domain.PessoaDto) error
	GetPersonById(personId string) (*domain.PessoaDto, error)
	SearchPerson(searchTerm string) ([]*domain.PessoaDto, error)
	Count() (int, error)
}
