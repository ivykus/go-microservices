package main

import (
	"log"
	"net/http"
)

const port = ":8080"

func main() {
	app := Config{}

	log.Println("Starting broker service on port", port)

	srv := &http.Server{
		Addr:    port,
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}

}
