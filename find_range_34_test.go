package roman_number

import (
	"reflect"
	"testing"
)

func Test_searchRange(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Example 1", args{[]int{5, 7, 7, 8, 8, 10}, 8}, []int{3, 4}},
		{"Submit failed 1", args{[]int{1}, 1}, []int{0, 0}},
		{"Submit failed 2", args{[]int{5, 7, 7, 8, 8, 10}, 8}, []int{3, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchRange(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("searchRange() = %v, want %v", got, tt.want)
			}
		})
	}
}
