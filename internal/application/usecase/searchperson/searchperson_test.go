package searchperson

import (
	"errors"
	"testing"

	"github.com/lucas-code42/rinha-backend/internal/domain"
	"github.com/lucas-code42/rinha-backend/tests/mock"
	"gotest.tools/assert"
)

var errorSearchPerson error = errors.New("error search person")

func TestSearchPersonUseCase(t *testing.T) {
	tableTest := []struct {
		name   string
		setup  func(mockRepository *mock.MockRepository)
		expect error
	}{
		{
			name: "success",
			setup: func(mockRepository *mock.MockRepository) {
				mockRepository.SearchPersonFunc = func(s string) ([]*domain.PessoaDto, error) {
					return []*domain.PessoaDto{
						{
							Id:         "123",
							Apelido:    "Jhon",
							Nome:       "Jhon Doe",
							Nascimento: "28/06/1914",
							Stack:      "golang;",
						},
						{
							Id:         "456",
							Apelido:    "Hom",
							Nome:       "Hommer Simpson",
							Nascimento: "12/05/1956",
							Stack:      "php;js;cobol;foxpro",
						},
					}, nil
				}
			},
			expect: nil,
		},
		{
			name: "empty list",
			setup: func(mockRepository *mock.MockRepository) {
				mockRepository.SearchPersonFunc = func(s string) ([]*domain.PessoaDto, error) {
					return []*domain.PessoaDto{}, nil
				}
			},
			expect: nil,
		},
		{
			name: "fail",
			setup: func(mockRepository *mock.MockRepository) {
				mockRepository.SearchPersonFunc = func(s string) ([]*domain.PessoaDto, error) {
					return []*domain.PessoaDto{}, errorSearchPerson
				}
			},
			expect: errorSearchPerson,
		},
	}

	for _, tt := range tableTest {
		t.Run(tt.name, func(t *testing.T) {
			mock := mock.NewMockRepository()
			tt.setup(mock)

			useCase := New(mock)
			_, result := useCase.Execute("foo")

			assert.Equal(t, tt.expect, result)
		})
	}
}
