package roman_number

import (
	"reflect"
	"testing"
)

func Test_kWeakestRows(t *testing.T) {
	type args struct {
		mat [][]int
		k   int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example 1",
			args: args{
				mat: [][]int{
					{1, 1, 0, 0, 0},
					{1, 1, 1, 1, 0},
					{1, 0, 0, 0, 0},
					{1, 1, 0, 0, 0},
					{1, 1, 1, 1, 1}},
				k: 3,
			},
			want: []int{2, 0, 3},
		},
		{
			name: "Submit failed #1",
			args: args{
				mat: [][]int{
					{1, 1, 0},
					{1, 1, 0},
					{1, 1, 1},
					{1, 1, 1},
					{0, 0, 0},
					{1, 1, 1},
					{1, 0, 0},
				},
				k: 6,
			},
			want: []int{4, 6, 0, 1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kWeakestRows(tt.args.mat, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("kWeakestRows() = %v, want %v", got, tt.want)
			}
		})
	}
}
