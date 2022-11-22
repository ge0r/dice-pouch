package rolls

import (
	"errors"
	"strconv"
	"strings"
)

type Roll struct {
	Name     string
	Expr     string
	Summands []Summand
	Sum      int
}

var ErrInvalidExpr = errors.New("roll: invalid expression")
var ErrNoSummands = errors.New("roll: no summands")

func New(name string, expr string) *Roll {
	return &Roll{
		Name: name,
		Expr: expr,
	}
}

func (r *Roll) Parse() error {
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
			if num == "" {
				continue
			}

			isDie := strings.Contains(num, "d")
			params := strings.Split(num, "d")

			var multiplier, val int
			var err1, err2 error

			if len(params) > 2 {
				return ErrInvalidExpr
			} else if len(params) > 1 {
				if params[0] != "" {
					// The first parameter is a multiplier and the second is a die
					multiplier, err1 = strconv.Atoi(params[0])
					val, err2 = strconv.Atoi(params[1])
				} else {
					// There is no multiplier and the second parameter is a die
					multiplier = 1
					val, err2 = strconv.Atoi(params[1])
				}
			} else {
				// There is only one parameter and it is a modifier
				multiplier = 1
				val, err2 = strconv.Atoi(params[0])
			}
			if err1 != nil || err2 != nil {
				return ErrInvalidExpr
			}
			for i := 0; i < multiplier; i++ {
				if isDie {
					res = append(res, &Dice{Sides: val, IsNegative: isNegative})
				} else {
					res = append(res, &Modifier{Value: val, IsNegative: isNegative})
				}
			}
			// All nums after the initial one are negative
			isNegative = true
		}
	}

	r.Summands = res
	return nil
}

func (r *Roll) roll() error {
	if len(r.Summands) == 0 {
		return ErrNoSummands
	}

	for _, summand := range r.Summands {
		r.Sum += summand.YieldRes()
	}

	return nil
}
