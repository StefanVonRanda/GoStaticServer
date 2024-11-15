package main

import (
	"log"
	"net/http"
	"strings"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		urlPath := r.URL.Path

		if strings.HasSuffix(urlPath, "/") {
			urlPath += "index.html"
		} else if !strings.Contains(urlPath, ".") {
			urlPath += ".html"
		}

		http.ServeFile(w, r, "./dist"+urlPath)
	})

	log.Println("Listening on localhost:3000...")
	if err := http.ListenAndServe(":3000", mux); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
