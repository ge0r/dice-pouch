package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		// TODO send a 404 and also display a message
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Welcome to dice pouch"))
}

func roll(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("No dice yet"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/roll", roll)

	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
