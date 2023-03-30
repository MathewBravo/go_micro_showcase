package main

import (
	"log"
	"net/http"
)

func (app *Config) Gateway(w http.ResponseWriter, r *http.Request) {
	payload := jsonReponse{
		Error:   false,
		Message: "Gateway recieved message",
	}

	err := app.jsonWrite(w, http.StatusOK, payload)
	if err != nil {
		log.Panicf("Could not write Json")
	}
}
