package controller

import (
	"backend-face/internal/entity"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type UserUseCase interface {
	RecognizeUser(user entity.User) (entity.User, error)
}

type userGetRequest struct {
	Encoding []float64 `json:"encoding"`
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
	if err := json.NewDecoder(request.Body).Decode(&userGet); err != nil || len(userGet.Encoding) != 256 {
		http.Error(writer, "wrong format", http.StatusBadRequest)
		return
	}

	var userEncoding entity.EncodingType
	copy(userEncoding[:256], userGet.Encoding)
	neededUser, err := controller.useCase.RecognizeUser(entity.User{Encoding: userEncoding})
	if err != nil {
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
