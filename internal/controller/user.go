package controller

import (
	"backend-face/internal/entity"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type UserUseCase interface {
	RecognizeUser(user entity.User) *entity.User
}

type userGetRequest struct {
	Encoding entity.EncodingType `json:"encoding"`
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

func (controller UserController) RecognizeUser(writer http.ResponseWriter, request *http.Request) {
	if method := request.Method; method != "POST" {
		return
	}

	var userGet userGetRequest
	if err := json.NewDecoder(request.Body).Decode(&userGet); err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	neededUser := controller.useCase.RecognizeUser(entity.User{Encoding: userGet.Encoding})
	if neededUser == nil {
		http.NotFound(writer, request)
		return
	}

	response, err := json.Marshal(userResponse{ID: neededUser.ID})
	if err != nil {
		log.Fatalln(err)
	}

	writer.Header().Set("Content-Type", "application/json")

	_, err = fmt.Fprint(writer, string(response))
	if err != nil {
		log.Fatalln(err)
	}
}

func (controller UserController) RegisterHandlers() {
	http.HandleFunc("/recognize", controller.RecognizeUser)
}
