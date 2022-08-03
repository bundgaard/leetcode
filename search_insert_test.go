package roman_number

import "testing"

func Test_searchInsert(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{[]int{1, 3, 5, 6}, 5}, 2},
		{"Example 2", args{[]int{1, 3, 5, 6}, 2}, 1},
		{"Example 3", args{[]int{1, 3, 5, 6}, 7}, 4},
		{"Submit failed 1", args{[]int{1, 3, 5, 6}, 0}, 0},
		{"Submit failed 2", args{[]int{2, 3, 4, 8, 10}, 0}, 0},
		{"Submit failed 3", args{[]int{3, 6, 7, 8, 10}, 5}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchInsert(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("searchInsert() = %v, want %v", got, tt.want)
			}
		})
	}
}
