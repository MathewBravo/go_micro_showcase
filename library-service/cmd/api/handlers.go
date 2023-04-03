package main

import (
	"encoding/json"
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

func (app *Config) BookById(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	book, err := app.Models.Book.GetBookByID(id)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	app.writeJSON(w, http.StatusAccepted, book)
}

func (app *Config) AddBook(w http.ResponseWriter, r *http.Request) {
	// `insert into books (title, authors, description, isreading, hasRead, thumbnail)
	// 			values($1,$2,$3,$4,$5)
	// 			return id
	// `

	decoder := json.NewDecoder(r.Body)
	book := app.Models.Book
	err := decoder.Decode(&book)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	_, err = app.Models.Book.Insert(book)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	app.writeJSON(w, http.StatusAccepted, "Added Book to Library")
}
