package p2918

import "testing"

func Test_minSum(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"Example 1", args{nums1: []int{3, 2, 0, 1, 0}, nums2: []int{6, 5, 0}}, 12},
		{"Example 2", args{nums1: []int{2, 0, 2, 0}, nums2: []int{1, 4}}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSum(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("minSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
