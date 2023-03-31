package main

import (
	"encoding/json"
	"io/ioutil"
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

type SearchTest struct {
	ItemCount int64 `json:"ItemCount"`
}

func (app *Config) BookSearch(w http.ResponseWriter, r *http.Request) {
	payload := jsonReponse{
		Error:   false,
		Message: "Messaged passed to search service",
	}

	resp, err := http.Get("http://go_micro_showcase_search-service_1/searchtest")
	if err != nil {
		payload.Error = true
		payload.Message = "Error fetching test results"
		payload.Data = err
		_ = app.jsonWrite(w, http.StatusNotFound, payload)
		return
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	st := SearchTest{}
	err = json.Unmarshal(respBody, &st)
	if err != nil {
		log.Println(err)
	}

	log.Println(st)
	payload.Data = st

	_ = app.jsonWrite(w, http.StatusOK, payload)

}
