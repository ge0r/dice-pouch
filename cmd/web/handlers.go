package main

import (
	"encoding/json"
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
	rollExpr := r.URL.Query()
	fmt.Println(rollExpr)

	rollsRes := []rolls.Roll{}

	if len(rollExpr) == 0 {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	for rollID, expr := range rollExpr {
		// Cannot have the same rollID twice
		if len(expr) > 1 {
			app.clientError(w, http.StatusBadRequest)
			return
		}

		roll := rolls.New(rollID, expr[0])
		err := roll.Execute()
		if err != nil {
			app.clientError(w, http.StatusBadRequest)
			return
		}
		rollsRes = append(rollsRes, *roll)
	}

	b, err := json.Marshal(rollsRes)
	if err != nil {
		app.serverError(w, err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}
