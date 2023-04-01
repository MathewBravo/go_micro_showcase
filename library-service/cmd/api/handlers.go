package main

import (
	"net/http"
)

func (app *Config) Library(w http.ResponseWriter, r *http.Request) {

	books, err := app.Models.Book.GetAll()
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	app.writeJSON(w, http.StatusAccepted, books)
}
