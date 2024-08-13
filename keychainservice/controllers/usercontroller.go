package controllers

import (
	"keychainservice/services"
)

type UserController struct {
	Service *services.UserService
}

func NewUserController(service *services.UserService) *UserController {
	return &UserController{Service: service}
}

func (uc *UserController) CreateUser(requestbody map[string]interface{}) map[string]interface{} {

	// For simplicity, just return the received user data
	response := map[string]interface{}{
		"message": "User created",
		"user":    "user_name",
	}
	return response

}
