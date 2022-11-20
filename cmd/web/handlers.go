package main

import (
	"fmt"
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}
	w.Write([]byte("Welcome to dice pouch"))
}

func (app *application) roll(w http.ResponseWriter, r *http.Request) {
	diceRoll := r.URL.Query()
	fmt.Println(diceRoll)
	// if diceRoll == "" {
	// 	app.clientError(w, http.StatusBadRequest)
	// 	return
	// }
	fmt.Fprintf(w, "No dice yet, but here is your dice roll: %s", diceRoll)
}
