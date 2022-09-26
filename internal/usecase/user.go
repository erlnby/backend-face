package usecase

import "backend-face/internal/entity"

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
	/*
		Здесь в user будет encoding. Нужно воспользоваться useCase.repository.GetAll, чтобы выбрать всех юзеров. Потом
		подобрать лучшее совпадение и вернуть.
	*/
	return entity.User{}
}
