package roman_number

import "testing"

func TestRoman(t *testing.T) {

	tests := []struct {
		Input    string
		Expected int
	}{
		{"III", 3},
		{"LVIII", 58},
		{"MCMXCIV", 1994},
	}

	for _, test := range tests {
		if got := romanToInt(test.Input); got != test.Expected {
			t.Errorf("expected %d. got %d\n", test.Expected, got)
		}
	}

}
