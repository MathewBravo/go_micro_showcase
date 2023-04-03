package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/lib/pq"
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

type Books struct {
	ID          int64          `json:"ID"`
	Title       string         `json:"title"`
	Authors     pq.StringArray `json:"authors"`
	Description string         `json:"despcription"`
	IsReading   bool           `json:"isReading"`
	HasRead     bool           `json:"hasRead"`
	Thumbnail   string         `json:"thumbnail"`
}

func (app *Config) AllLibrary(w http.ResponseWriter, r *http.Request) {
	payload := jsonReponse{
		Error:   false,
		Message: "Message passed to library service",
	}

	resp, err := http.Get("http://go_micro_showcase_library-service_1/library")
	if err != nil {
		payload.Error = true
		payload.Message = "Error fetching all books"
		payload.Data = err
		_ = app.jsonWrite(w, http.StatusNotFound, payload)
		return
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	log.Println(respBody)

	b := []Books{}
	err = json.Unmarshal(respBody, &b)
	if err != nil {
		log.Println(err)
	}

	log.Println(b)
	payload.Data = b

	_ = app.jsonWrite(w, http.StatusOK, payload)
}

func (app *Config) AddBook(w http.ResponseWriter, r *http.Request) {

}
