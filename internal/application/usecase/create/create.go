package create

import (
	"github.com/lucas-code42/rinha-backend/internal/application"
	"github.com/lucas-code42/rinha-backend/internal/domain"
)

type CreateUseCase struct {
	pessoa     *domain.PessoaDto
	repository application.RespositoryImpl
}

func New(
	pessoa *domain.PessoaDto,
	repository application.RespositoryImpl,
) *CreateUseCase {
	return &CreateUseCase{
		pessoa:     pessoa,
		repository: repository,
	}
}

func (c *CreateUseCase) Execute() error {
	return c.repository.CreatePerson(c.pessoa)
}
