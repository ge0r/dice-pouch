package rolls

import (
	"reflect"
	"testing"
)

type testData struct {
	input    string
	expected []string
	err      error
}

func TestParse(t *testing.T) {
	tests := []struct {
		name    string
		expr    string
		want    []Summand
		wantErr bool
	}{
		{
			"Single dice with modifier",
			"d20 5",
			[]Summand{
				&Dice{20, false},
				&Modifier{5, false},
			},
			false,
		},
		{
			"Two different dice with multipliers and a plus modifier",
			"3d20+15 2d6",
			[]Summand{
				&Dice{20, false},
				&Dice{20, false},
				&Dice{20, false},
				&Modifier{15, false},
				&Dice{6, false},
				&Dice{6, false},
			},
			false,
		},
		{
			"Negative modifier",
			"d12 15 2d4-3",
			[]Summand{
				&Dice{12, false},
				&Modifier{15, false},
				&Dice{4, false},
				&Dice{4, false},
				&Modifier{3, true},
			},
			false,
		},
		{
			"Negative dice and two negative modifiers",
			"-d1 -11 d3-2d4-6",
			[]Summand{
				&Dice{1, true},
				&Modifier{11, true},
				&Dice{3, false},
				&Dice{4, true},
				&Dice{4, true},
				&Modifier{6, true},
			},
			false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			roll := New("test", tt.expr)
			res, err := roll.Parse()
			// try if (err==nil) == tt.wantErr
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
			}
			a := len(res)
			b := len(tt.want)
			if a != b {
				t.Errorf("Parse() result size = %d, wants %d", len(res), len(tt.want))
				return
			}
			for i, val := range tt.want {
				if !reflect.DeepEqual(res[i], val) {
					t.Errorf("Parse() result[%d] = %v, wants %v", i, res[i], val)
				}
			}
		})
	}
}
