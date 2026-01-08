package p1458

import "testing"

func Test_maxDotProduct(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{nums1: []int{2, 1, -2, 5}, nums2: []int{3, 0, -6}}, 18},
		{"Example 2", args{nums1: []int{3, -2}, nums2: []int{2, -6, 7}}, 21},
		{"Example 3", args{nums1: []int{-1, -1}, nums2: []int{1, 1}}, -1},
		{"TestCase 11", args{nums1: []int{5, -4, -3}, nums2: []int{-4, -3, 0, -4, 2}}, 28},
		{"TestCase 39", args{nums1: []int{7, 2, 2, -1, -1, 1, -4, 7, 6}, nums2: []int{-7, -9, -1, 2, 2, 5, -7, 2, -7, 5}}, 108},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDotProduct(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("maxDotProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
