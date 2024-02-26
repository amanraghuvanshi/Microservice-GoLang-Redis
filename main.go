package main

import (
	"log"
	"net/http"
)

func main() {
	// Initializing the server
	server := &http.Server{
		Addr:    ":3000",
		Handler: http.HandlerFunc(basicHandler),
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatalln("Failed to start the server", err)
	}
}

func basicHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!!"))
}
