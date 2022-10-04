package usecase

import (
	"backend-face/internal/entity"
	"backend-face/internal/utils"
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

	minScore := math.MaxFloat64
	var neededUser entity.User

	for _, userDB := range users {
		distance, _ := utils.GetEuclideanDistance(userDB.Encoding[:], user.Encoding[:])
		if distance <= Tolerance && distance < minScore {
			minScore = distance
			neededUser = userDB
		}
	}

	return neededUser
}
