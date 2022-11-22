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
				&Dice{
					Sides:      20,
					IsNegative: false,
				},
				&Modifier{
					Value:      5,
					IsNegative: false,
				},
			},
			false,
		},
		{
			"Two different dice with multipliers and a plus modifier",
			"3d20+15 2d6",
			[]Summand{
				&Dice{
					Sides:      20,
					IsNegative: false,
				},
				&Dice{
					Sides:      20,
					IsNegative: false,
				},
				&Dice{
					Sides:      20,
					IsNegative: false,
				},
				&Modifier{
					Value:      15,
					IsNegative: false,
				},
				&Dice{
					Sides:      6,
					IsNegative: false,
				},
				&Dice{
					Sides:      6,
					IsNegative: false,
				},
			},
			false,
		},
		{
			"Negative modifier",
			"d12 15 2d4-3",
			[]Summand{
				&Dice{
					Sides:      12,
					IsNegative: false,
				},
				&Modifier{
					Value:      15,
					IsNegative: false,
				},
				&Dice{
					Sides:      4,
					IsNegative: false,
				},
				&Dice{
					Sides:      4,
					IsNegative: false,
				},
				&Modifier{
					Value:      3,
					IsNegative: true,
				},
			},
			false,
		},
		{
			"Negative dice and two negative modifiers",
			"-d1 -11 d3-2d4-6",
			[]Summand{
				&Dice{
					Sides:      1,
					IsNegative: true,
				},
				&Modifier{
					Value:      11,
					IsNegative: true,
				},
				&Dice{
					Sides:      3,
					IsNegative: false,
				},
				&Dice{
					Sides:      4,
					IsNegative: true,
				},
				&Dice{
					Sides:      4,
					IsNegative: true,
				},
				&Modifier{
					Value:      6,
					IsNegative: true,
				},
			},
			false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			roll := New("test", tt.expr)
			err := roll.Parse()
			// try if (err==nil) == tt.wantErr
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			res := roll.Summands
			if len(res) != len(tt.want) {
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
