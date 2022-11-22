package main

import (
	"fmt"
	"net/http"

	"github.com/ge0r/dice-pouch/pkg/rolls"
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

	if len(diceRoll) == 0 {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	for rollID, expr := range diceRoll {
		// Cannot have the same rollID twice
		if len(expr) > 1 {
			app.clientError(w, 400)
		}

		roll := rolls.New(rollID, expr[0])
		err := roll.Parse()
		err = roll.Roll()
		fmt.Println(err)
		fmt.Println(roll.Sum)
	}
	fmt.Fprintf(w, "No dice yet, but here is your dice roll: %s", diceRoll)
}
