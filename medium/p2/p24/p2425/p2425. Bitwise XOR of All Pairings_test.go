package p2425

import "testing"

func Test_xorAllNums(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{nums1: []int{2, 1, 3}, nums2: []int{10, 2, 5, 0}}, 13},
		{"Example 2", args{nums1: []int{1, 2}, nums2: []int{3, 4}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := xorAllNums(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("xorAllNums() = %v, want %v", got, tt.want)
			}
		})
	}
}
