package rolls

import "math/rand"

type Summand interface {
	YieldRes() int
}

type Dice struct {
	Sides      int
	IsNegative bool
	Result     int
}

type Modifier struct {
	Value      int
	IsNegative bool
}

func (d *Dice) YieldRes() int {
	// Only calculate result the first time
	if d.Result == 0 {
		res := rand.Int()%d.Sides + 1

		if d.IsNegative {
			d.Result = -res
		} else {
			d.Result = res
		}
	}

	return d.Result
}

func (m *Modifier) YieldRes() int {
	if m.IsNegative {
		return -m.Value
	} else {
		return m.Value
	}
}
