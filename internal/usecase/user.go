package usecase

import (
	"backend-face/internal/entity"
	"backend-face/internal/utils"
	"errors"
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

func (useCase UserUseCase) RecognizeUser(user entity.User) (entity.User, error) {
	users := useCase.repository.GetAll()

	var neededUser entity.User
	err := errors.New("User not found")
	minScore := math.MaxFloat64

	for _, userDB := range users {
		distance, _ := utils.GetEuclideanDistance(userDB.Encoding[:], user.Encoding[:])
		if distance <= Tolerance && distance < minScore {
			err = nil
			minScore = distance
			neededUser = userDB
		}
	}

	return neededUser, err
}
