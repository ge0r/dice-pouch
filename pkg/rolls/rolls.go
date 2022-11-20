package rolls

import "errors"

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

// write unit test
// 3d20 15 d6 || 3d20 15 d6-3
func (r *Roll) Parse() ([]string, error) {
	return nil, nil
}

// func (r *Roll) execute() map[string]int {
// }
