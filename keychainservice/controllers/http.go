package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func NewHTTPController(httpMethod string, controller func(interface{}) map[string]interface{}, dto interface{}) func(w http.ResponseWriter, r *http.Request) {
	// Validations for request body

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != httpMethod {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		requestBodyBytes, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// var requestBody map[string]interface{}
		dto2 := UserStructBody{}

		err = json.Unmarshal(requestBodyBytes, &dto2)
		if err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}
		fmt.Println("DTO", dto2)

		// validate := validator.New()

		response := controller(dto2)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)
	}

}
