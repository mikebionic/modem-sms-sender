package main

import (
	"encoding/json"
	"net/http"
)

type BasicApiResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func respondWithJSON(w http.ResponseWriter, code int, payload map[string]interface{}) {

	var apiResponse BasicApiResponse

	if !isNil(payload["status"]) {
		apiResponse.Status = payload["status"].(int)
	} else {
		apiResponse.Status = 0
	}

	if !isNil(payload["message"].(string)) {
		apiResponse.Message = payload["message"].(string)
	}

	response, _ := json.Marshal(apiResponse)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]interface{}{"message": message})
}
