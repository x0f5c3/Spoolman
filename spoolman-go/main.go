package main

import (
	"log"
	"net/http"

	// Generated API handlers from OpenAPI specification.
	_ "spoolman-go/api"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Spoolman Go prototype"))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
