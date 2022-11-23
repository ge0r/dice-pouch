package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"runtime/debug"
	"strconv"

	"github.com/ge0r/dice-pouch/pkg/rolls"
)

type ResultJson struct {
	rollArray [][]map[string]int
}

func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Output(2, trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *application) notFound(w http.ResponseWriter) {
	http.Error(w, "Not all those who wander are lost", http.StatusNotFound)
}

func (app *application) resultJson(results []rolls.Roll) ([]byte, error) {
	res := [][]map[string]int{}

	for _, result := range results {
		var label string

		r := []map[string]int{}
		// Append the sum of the roll first
		r = append(r, map[string]int{"sum": result.Sum})
		for _, summand := range result.Summands {
			result := summand.YieldRes()
			switch s := summand.(type) {
			case *rolls.Dice:
				// This is a die, map the d20 label to the result
				label = "d" + strconv.Itoa(s.Sides)
			case *rolls.Modifier:
				// This is a modifier
				label = "modifier"
			}
			// Add the label:result key pair to the array
			r = append(r, map[string]int{label: result})
		}
		res = append(res, r)
	}

	b, err := json.Marshal(res)
	if err != nil {
		return nil, err
	}
	return b, nil
}
