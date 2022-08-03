package roman_number

import "testing"

func TestRansomNote(t *testing.T) {
	tests := []struct {
		RansomNote string
		Magazine   string
		Expected   bool
	}{
		{
			RansomNote: "a",
			Magazine:   "b",
			Expected:   false,
		},
		{
			RansomNote: "aa",
			Magazine:   "ab",
			Expected:   false,
		},
		{
			RansomNote: "aa",
			Magazine:   "aab",
			Expected:   true,
		},
	}

	for _, test := range tests {
		if got := canConstruct(test.RansomNote, test.Magazine); got != test.Expected {
			t.Logf("expected %v. Got %v\n", test.Expected, got)
		}
	}
}
