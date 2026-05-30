package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "ok",
		"version": "1.0.2",
	})
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "world"
	}
	json.NewEncoder(w).Encode(map[string]string{
		"message": "hello, " + name,
	})
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/hello", helloHandler)

	log.Printf("server listening on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
