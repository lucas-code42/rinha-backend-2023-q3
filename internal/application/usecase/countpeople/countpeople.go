package countpeople

import (
	"github.com/lucas-code42/rinha-backend/internal/application"
)

type CountPeople struct {
	repository application.RespositoryImpl
}

func New(
	repository application.RespositoryImpl,
) *CountPeople {
	return &CountPeople{
		repository: repository,
	}
}

func (c *CountPeople) Execute() (int, error) {
	return c.repository.Count()
}
