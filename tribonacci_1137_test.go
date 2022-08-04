package roman_number

import "testing"

func Test_tribonacci(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{4}, 4}, // TODO: Add test cases.
		{"Example 2", args{25}, 1389537},
		{"Submit failed 1", args{35}, 615693474},
		{"Submit failed 2", args{0}, 0},
		{"Submit failed 3", args{1}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tribonacci(tt.args.n); got != tt.want {
				t.Errorf("tribonacci() = %v, want %v", got, tt.want)
			}
		})
	}
}
