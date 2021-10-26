package main

import (
	"log"
	"net/http"
	"time"
)

func main() {

	server := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  30 * time.Second,
	}

	staticDirectory := http.FileServer(http.Dir("./static"))

	http.Handle("/", staticDirectory)

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
