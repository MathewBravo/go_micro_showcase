package main

import (
	"fmt"
	"log"
	"net/http"
)

const PORT = "80"

type Config struct {
}

func main() {
	log.Println("Starting Library Service")

	app := Config{}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", PORT),
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
