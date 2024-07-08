package getpersonbyid

import (
	"errors"
	"fmt"
	"testing"

	"github.com/lucas-code42/rinha-backend/internal/domain"
	"github.com/lucas-code42/rinha-backend/tests/mock"
	"gotest.tools/assert"
)

var errorGetPersonById error = errors.New("error get person by id")

func TestGetPersonByIdUseCase(t *testing.T) {
	tableTests := []struct {
		name     string
		setup    func(mockRepository *mock.MockRepository)
		expected error
	}{
		{
			name: "success",
			setup: func(mockRepository *mock.MockRepository) {
				mockRepository.GetPersonByIdFunc = func(s string) (*domain.PessoaDto, error) {
					return &domain.PessoaDto{
						Id:         "123",
						Apelido:    "Jhon",
						Nome:       "Jhon Doe",
						Nascimento: "09/11/1989",
						Stack:      "c;c++;c#",
					}, nil
				}
			},
			expected: nil,
		},
		{
			name: "fail",
			setup: func(mockRepository *mock.MockRepository) {
				mockRepository.GetPersonByIdFunc = func(s string) (*domain.PessoaDto, error) {
					return &domain.PessoaDto{}, errorGetPersonById
				}
			},
			expected: errorGetPersonById,
		},
	}

	for _, tt := range tableTests {
		t.Run(tt.name, func(t *testing.T) {
			mock := mock.NewMockRepository()
			tt.setup(mock)

			useCase := New(mock)
			_, result := useCase.Execute("123")

			assert.Equal(
				t,
				tt.expected,
				result, fmt.Sprintf("name: '%s', expected: '%v' but got '%v'", tt.name, tt.expected, result),
			)
		})
	}

}
