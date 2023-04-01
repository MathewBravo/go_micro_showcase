package main

import (
	"fmt"
	"net/http"
)

func (app *Config) Library(w http.ResponseWriter, r *http.Request) {

	books, err := app.Models.Book.GetAll()
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprint("Found the attached books"),
		Data:    books,
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}
