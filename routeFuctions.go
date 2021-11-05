package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type RequestData struct {
	PhoneNumber string `json:"phone_number"`
	MessageText string `json:"message_text"`
	Token       string `json:"token"`
}

func (a *App) apiSendSMS(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}

	var sms_request_data RequestData
	json.Unmarshal(reqBody, &sms_request_data)

	sha_key, err := get_sha_key_from_config()
	if err != nil {
		return
	}
	if sms_request_data.Token != sha_key {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}

	err = send_sms(sms_request_data.PhoneNumber, sms_request_data.MessageText)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	res := make(map[string]interface{})
	res["message"] = "success"
	res["status"] = 1

	respondWithJSON(w, http.StatusOK, res)
}

func (a *App) getRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "SMS sender API"}`))
}

func (a *App) notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"message": "Not found"}`))
}
