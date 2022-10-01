package controller

import (
	"backend-face/internal/entity"
	"encoding/json"
	"fmt"
	"log"
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
	if method := r.Method; method != "POST" {
		return
	}
	var b userGetRequest
	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	neededUser := controller.useCase.RecognizeUser(entity.User{Encoding: b.Encoding})
	if neededUser.ID == "" {
		http.NotFound(w, r)
		return
	}
	result, err := json.Marshal(userResponse{ID: neededUser.ID})
	if err != nil {
		log.Fatalln(err)
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = fmt.Fprint(w, string(result))
	if err != nil {
		log.Fatalln(err)
	}

}

func (controller UserController) RegisterHandlers() {
	http.HandleFunc("/recognize", controller.RecognizeUser)
}
