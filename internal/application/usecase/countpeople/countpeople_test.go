package countpeople

import (
	"errors"
	"testing"

	"github.com/lucas-code42/rinha-backend/tests/mock"
	"gotest.tools/assert"
)

var errorCount error = errors.New("count error")

func TestCountPeopleUsecase(t *testing.T) {
	tableTests := []struct {
		name     string
		setup    func(mockRepository *mock.MockRepository)
		expected error
	}{
		{
			name: "sucess",
			setup: func(mockRepository *mock.MockRepository) {
				mockRepository.CountFunc = func() (int, error) {
					return 10, nil
				}
			},
			expected: nil,
		},
		{
			name: "fail",
			setup: func(mockRepository *mock.MockRepository) {
				mockRepository.CountFunc = func() (int, error) {
					return 0, errorCount
				}
			},
			expected: errorCount,
		},
	}

	for _, tt := range tableTests {
		t.Run(tt.name, func(t *testing.T) {
			mock := mock.NewMockRepository()
			tt.setup(mock)

			useCase := New(mock)
			_, result := useCase.Execute()

			assert.Equal(t, tt.expected, result)
		})
	}
}
