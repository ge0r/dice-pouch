package rolls

import (
	"errors"
	"strconv"
	"strings"
)

type Roll struct {
	Name string
	Expr string
	Res  map[string]int
	Sum  int
}

var ErrInvalidExpr = errors.New("roll: invalid expression")

func New(name string, expr string) *Roll {
	return &Roll{
		Name: name,
		Expr: expr,
	}
}

func (r *Roll) Parse() ([]Summand, error) {
	// Parse() treats all spaces as "+"
	// So we need to make sure that there are no "+" left
	expr := strings.ReplaceAll(r.Expr, "+", " ")

	words := strings.Fields(expr)

	res := []Summand{}
	for _, word := range words {
		isNegative := false
		// If the first char is -, initial num is negative.
		if word[0:1] == "-" {
			isNegative = true
		}

		// If the word contains more negative nums, split them
		nums := strings.Split(word, "-")

		for _, num := range nums {
			isDie := strings.Contains(num, "d")
			params := strings.Split(num, "d")

			var multiplier, val int
			var err1, err2 error

			if len(params) > 2 {
				return nil, ErrInvalidExpr
			} else if len(params) > 1 {
				// The first parameter is a multiplier and the second is a die
				multiplier, err1 = strconv.Atoi(params[0])
				val, err2 = strconv.Atoi(params[1])
				if err1 != nil || err2 != nil {
					return nil, ErrInvalidExpr
				}
			} else {
				// There is only one parameter, either modifier or die
				multiplier = 1
				val, err2 = strconv.Atoi(params[0])
			}
			for i := 0; i < multiplier; i++ {
				if isDie {
					res = append(res, &Dice{val, isNegative})
				} else {
					res = append(res, &Modifier{val, isNegative})
				}
			}
			// All nums after the initial one are negative
			isNegative = true
		}
	}

	return res, nil
}
