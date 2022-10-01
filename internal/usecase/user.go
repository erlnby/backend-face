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
		sum        float64 = 0
		minScore   float64 = math.MaxFloat64
		neededUser entity.User
	)
	for i := 0; i < len(users); i++ {
		for j := 0; j < 256; j++ {
			sum += math.Pow(users[i].Encoding[j]-user.Encoding[j], 2)
		}
		if res := math.Sqrt(sum); res <= Tolerance && res < minScore {
			minScore = res
			neededUser = users[i]
		}
		sum = 0
	}
	return neededUser
}
