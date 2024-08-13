package controllers

import (
	"fmt"
	"keychainservice/services"
)

type UserController struct {
	Service *services.UserService
}

func NewUserController(service *services.UserService) *UserController {
	return &UserController{Service: service}
}

type UserStructBody struct {
	UserName string `json:"username" validate:"required"`
	Age      int    `json:"age" validate:"gte=18,lte=100"`
}

func (uc *UserController) CreateUser(requestbody interface{}) map[string]interface{} {

	// For simplicity, just return the received user data
	fmt.Print("<><>", requestbody)
	// map: key value --> X struct  --> X user struct

	user, ok := requestbody.(*UserStructBody)

	if !ok {
		fmt.Println("error")
		return nil
	}

	response := map[string]interface{}{
		"message": "User created",
		"user":    (*user).UserName,
		"age":     (*user).Age,
	}
	return response

}
