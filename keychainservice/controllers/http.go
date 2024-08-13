package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
)

func NewHTTPController(
	httpMethod string,
	controller func(interface{}) map[string]interface{},
	dtoType reflect.Type,
	middlewareList []func(w http.ResponseWriter, r *http.Request),
) func(w http.ResponseWriter, r *http.Request) {
	// Validations for request body

	return func(w http.ResponseWriter, r *http.Request) {
		for _, middleware := range middlewareList {
			middleware(w, r)
		}

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
		// fmt.Println(reflect.TypeOf(dto), "SOMETHING")
		dtoValue := reflect.New(dtoType).Interface()
		err = json.Unmarshal(requestBodyBytes, &dtoValue)
		if err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}
		fmt.Println("DTO", reflect.TypeOf(dtoValue))

		// validate := validator.New()

		response := controller(dtoValue)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)
	}

}
