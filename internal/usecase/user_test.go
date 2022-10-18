package usecase_test

import (
	"backend-face/internal/entity"
	"backend-face/internal/usecase"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"math"
	"math/rand"
	"testing"
)

type UserRepositoryMock struct {
	mock.Mock
}

func (repositoryMock *UserRepositoryMock) GetAll() []entity.User {
	return repositoryMock.Called().Get(0).([]entity.User)
}

func TestUserUseCase_RecognizeUser(t *testing.T) {
	t.Run("User can't be recognized if there are no users", func(t *testing.T) {
		var users []entity.User
		testUser := entity.User{}

		repositoryMock := UserRepositoryMock{}
		repositoryMock.On("GetAll").Return(users)
		useCase := usecase.NewUserUseCase(&repositoryMock)
		recognizedUser := useCase.RecognizeUser(testUser)

		assert.Nil(t, recognizedUser)
	})

	t.Run("User can't be recognized if a minimum score is higher than tolerance", func(t *testing.T) {
		var userEncodingExample entity.EncodingType
		for i := 0; i < len(userEncodingExample); i++ {
			userEncodingExample[i] = math.MaxFloat64
		}
		users := []entity.User{
			{ID: "1", Encoding: userEncodingExample},
		}
		testUser := entity.User{}

		repositoryMock := UserRepositoryMock{}
		repositoryMock.On("GetAll").Return(users)
		useCase := usecase.NewUserUseCase(&repositoryMock)
		recognizedUser := useCase.RecognizeUser(testUser)

		assert.Nil(t, recognizedUser)
	})

	t.Run("User must be recognized with a best match", func(t *testing.T) {
		var userEncodingExample entity.EncodingType
		for i := 0; i < len(userEncodingExample); i++ {
			userEncodingExample[i] = 1
		}
		users := []entity.User{
			{ID: "1", Encoding: userEncodingExample},
			{ID: "2"},
		}
		testUser := entity.User{}

		repositoryMock := UserRepositoryMock{}
		repositoryMock.On("GetAll").Return(users)
		useCase := usecase.NewUserUseCase(&repositoryMock)
		recognizedUser := useCase.RecognizeUser(testUser)

		assert.Equal(t, recognizedUser.ID, "2")
	})
}

func BenchmarkUserUseCase_RecognizeUser(b *testing.B) {
	const usersCount = 1000

	user := entity.User{}
	users := make([]entity.User, usersCount)

	for userID := 0; userID < usersCount; userID++ {
		var encoding entity.EncodingType
		for i := 0; i < len(encoding); i++ {
			encoding[i] = rand.NormFloat64()
		}
		users = append(users, entity.User{ID: fmt.Sprint(userID), Encoding: encoding})
	}

	repositoryMock := UserRepositoryMock{}
	repositoryMock.On("GetAll").Return(users)
	useCase := usecase.NewUserUseCase(&repositoryMock)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		useCase.RecognizeUser(user)
	}
}
