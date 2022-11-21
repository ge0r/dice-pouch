package rolls

import "math/rand"

type Summand interface {
	YieldNum() int
}

type Dice struct {
	Sides      int
	isNegative bool
}

type Modifier struct {
	Value      int
	isNegative bool
}

func (d *Dice) YieldNum() int {
	res := rand.Int()%d.Sides + 1

	if d.isNegative {
		return -res
	} else {
		return res
	}
}

func (m *Modifier) YieldNum() int {
	if m.isNegative {
		return -m.Value
	} else {
		return m.Value
	}
}
