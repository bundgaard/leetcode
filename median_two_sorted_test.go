package roman_number

import (
	"runtime"
	"testing"
	"time"
)

func Test_findMedianSortedArrays(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"Example 1", args{[]int{1, 3}, []int{2}}, 2.0},
		{"Example 2", args{[]int{1, 2}, []int{3, 4}}, 2.5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s1 := time.Now()

			if got := findMedianSortedArrays(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("findMedianSortedArrays() = %v, want %v", got, tt.want)
			}
			s2 := time.Since(s1)
			t.Logf("spent %s", s2)
			ms := runtime.MemStats{}
			runtime.ReadMemStats(&ms)

			t.Logf("Memory %v", ms.TotalAlloc)
		})
	}
}
