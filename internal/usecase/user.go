package usecase

import (
	"backend-face/internal/entity"
	"math"
)

const Tolerance = 0.6

type UserRepository interface {
	GetAll() []entity.User
}

type UserUseCase struct {
	repository UserRepository
}

func NewUserUseCase(repository UserRepository) UserUseCase {
	return UserUseCase{repository}
}

func (useCase UserUseCase) RecognizeUser(user entity.User) entity.User {
	users := useCase.repository.GetAll()

	var (
		minScore   float64 = math.MaxFloat64
		neededUser entity.User
	)

	for _, userDB := range users {
		sum := 0.0

		for j := 0; j < 256; j++ {
			sum += math.Pow(userDB.Encoding[j]-user.Encoding[j], 2)
		}

		if res := math.Sqrt(sum); res <= Tolerance && res < minScore {
			minScore = res
			neededUser = userDB
		}
	}

	return neededUser
}
