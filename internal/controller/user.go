package controller

import (
	"backend-face/internal/entity"
	"net/http"
)

type UserUseCase interface {
	RecognizeUser(user entity.User) entity.User
}

type userGetRequest struct {
	Encoding [256]float64 `json:"encoding"`
}

type userResponse struct {
	ID string `json:"user_id"`
}

type UserController struct {
	useCase UserUseCase
}

func NewUserController(useCase UserUseCase) UserController {
	return UserController{useCase}
}

func (controller UserController) RecognizeUser(w http.ResponseWriter, r *http.Request) {
	/*
		Нужно проверсти преобразование r -> json (по структуре userGetRequest) -> entity.User (неполностью заполненная).
		useCase вернет также entity.User. Ее нужно преобразовать в userResponse, потом в json и отдать в w. На всех
		преобразованиях предусмостреть HTTP-ошибки.
	*/
	controller.useCase.RecognizeUser(entity.User{})
}

func (controller UserController) RegisterHandlers() {
	http.HandleFunc("/recognize", controller.RecognizeUser)
}
