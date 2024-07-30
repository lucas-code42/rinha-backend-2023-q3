package createperson

import (
	"errors"
	"testing"

	"github.com/lucas-code42/rinha-backend/internal/domain"
	"github.com/lucas-code42/rinha-backend/tests/mock"
	"gotest.tools/assert"
)

var errorCreatePerson error = errors.New("error create person")

func TestCreatePersonUseCase(t *testing.T) {
	tableTests := []struct {
		name      string
		setup     func(mockRepository *mock.MockRepository)
		personDto *domain.PessoaDto
		expected  error
	}{
		{
			name: "Success",
			setup: func(mockRepository *mock.MockRepository) {
				mockRepository.CreatePersonFunc = func(pd *domain.PessoaDto) error {
					return nil
				}
			},
			personDto: &domain.PessoaDto{
				Id:         "",
				Apelido:    "jhon",
				Nome:       "jhon Doe",
				Nascimento: "02/09/1945",
				Stack:      "golang;python;rust",
			},
			expected: nil,
		},
		{
			name: "fail",
			setup: func(mockRepository *mock.MockRepository) {
				mockRepository.CreatePersonFunc = func(pd *domain.PessoaDto) error {
					return errorCreatePerson
				}
			},
			personDto: &domain.PessoaDto{},
			expected:  errorCreatePerson,
		},
	}

	for _, tt := range tableTests {
		t.Run(tt.name, func(t *testing.T) {
			mock := mock.NewMockRepository()
			tt.setup(mock)

			useCase := New(tt.personDto, mock)
			result := useCase.Execute()

			assert.Equal(t, tt.expected, result)
		})
	}
}
