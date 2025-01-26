package main

import (
	"log"
	"net/http"
)

const webPort = "80"

type Config struct {
}

func main() {
	app := Config{}

	log.Printf("Starting server on port: %s", webPort)

	srv := &http.Server{
		Addr:    ":" + webPort,
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
