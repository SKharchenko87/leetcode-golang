package p1855

import "testing"

func Test_maxDistance(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{nums1: []int{55, 30, 5, 4, 2}, nums2: []int{100, 20, 10, 10, 5}}, 2},
		{"Example 2", args{nums1: []int{2, 2, 2}, nums2: []int{10, 10, 1}}, 1},
		{"Example 3", args{nums1: []int{30, 29, 19, 5}, nums2: []int{25, 25, 25, 25, 25}}, 2},
		{"TestCase 20", args{nums1: []int{5, 4}, nums2: []int{3, 2}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDistance(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("maxDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
