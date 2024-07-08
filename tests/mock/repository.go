package mock

import "github.com/lucas-code42/rinha-backend/internal/domain"

type MockRepository struct {
	CreatePersonFunc  func(pd *domain.PessoaDto) error
	GetPersonByIdFunc func(s string) (*domain.PessoaDto, error)
	SearchPersonFunc  func(s string) ([]*domain.PessoaDto, error)
	CountFunc         func() (int, error)
}

func (m *MockRepository) CreatePerson(person *domain.PessoaDto) error {
	return m.CreatePersonFunc(&domain.PessoaDto{})
}

func (m *MockRepository) GetPersonById(personId string) (*domain.PessoaDto, error) {
	return m.GetPersonByIdFunc("")
}

func (m *MockRepository) SearchPerson(searchTerm string) ([]*domain.PessoaDto, error) {
	return m.SearchPersonFunc("")
}

func (m *MockRepository) Count() (int, error) {
	return m.CountFunc()
}

func NewMockRepository() *MockRepository {
	return &MockRepository{
		CreatePersonFunc:  func(pd *domain.PessoaDto) error { return nil },
		GetPersonByIdFunc: func(s string) (*domain.PessoaDto, error) { return &domain.PessoaDto{}, nil },
		SearchPersonFunc:  func(s string) ([]*domain.PessoaDto, error) { return []*domain.PessoaDto{}, nil },
		CountFunc:         func() (int, error) { return 0, nil },
	}
}
