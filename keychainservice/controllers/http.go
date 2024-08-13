package controllers

import (
	"encoding/json"
	"net/http"
)

func NewHTTPController(httpMethod string, controller func(map[string]interface{}) map[string]interface{}) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != httpMethod {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		var requestBody map[string]interface{}
		if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		response := controller(requestBody)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)
	}

}
