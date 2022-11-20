package rolls

import "testing"

type testData struct {
	input    string
	expected []string
	err      error
}

func TestParse(t *testing.T) {
	tests := []struct {
		name    string
		expr    string
		want    []string
		wantErr bool
	}{
		{
			"Single dice with modifier",
			"d20 5",
			[]string{"d20", "5"},
			false,
		},
		{
			"Two different dice, one with multiplier, plus modifier",
			"3d20 15 d6",
			[]string{"d20", "d20", "d20", "15", "d6"},
			false,
		},
		{
			"Negative modifier",
			"d12 15 2d4-3",
			[]string{"d12", "15", "d4", "d4", "-3"},
			false,
		},
		{
			"Two negative modifiers",
			"d1 11 d3-2d4-6",
			[]string{"d1", "11", "d3", "-d4", "-d4", "-6"},
			false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			roll := New("test", tt.expr)
			result, err := roll.Parse()
			// try if (err==nil) == tt.wantErr
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
			}
			if len(result) != len(tt.want) {
				t.Errorf("Parse() result size = %d, wants %d", len(result), len(tt.want))
			}
			for i, val := range tt.want {
				if result[i] != val {
					t.Errorf("Parse() result[%d] = %s, wants %s", i, result[i], val)
				}
			}
		})
	}
}
