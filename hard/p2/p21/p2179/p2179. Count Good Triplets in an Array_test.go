package p2179

import "testing"

func Test_goodTriplets(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"Example 1", args{nums1: []int{2, 0, 1, 3}, nums2: []int{0, 1, 2, 3}}, 1},
		{"Example 2", args{nums1: []int{4, 0, 1, 3, 2}, nums2: []int{4, 1, 0, 2, 3}}, 4},
		{"TestCase 38", args{nums1: []int{13, 14, 10, 2, 12, 3, 9, 11, 15, 8, 4, 7, 0, 6, 5, 1}, nums2: []int{8, 7, 9, 5, 6, 14, 15, 10, 2, 11, 4, 13, 3, 12, 1, 0}}, 77},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := goodTriplets(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("goodTriplets() = %v, want %v", got, tt.want)
			}
		})
	}
}
