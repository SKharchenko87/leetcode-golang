package p2333

import "testing"

func Test_minSumSquareDiff(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
		k1    int
		k2    int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"Example 1", args{nums1: []int{1, 2, 3, 4}, nums2: []int{2, 10, 20, 19}, k1: 0, k2: 0}, 579},
		{"Example 2", args{nums1: []int{1, 4, 10, 12}, nums2: []int{5, 8, 6, 9}, k1: 1, k2: 1}, 43},
		{"TestCase 2", args{nums1: []int{1, 4, 10, 12}, nums2: []int{5, 8, 6, 9}, k1: 10, k2: 5}, 0},
		{"TestCase 7", args{nums1: []int{7, 5, 0, 12, 14}, nums2: []int{7, 5, 0, 12, 14}, k1: 2, k2: 9}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSumSquareDiff(tt.args.nums1, tt.args.nums2, tt.args.k1, tt.args.k2); got != tt.want {
				t.Errorf("minSumSquareDiff() = %v, want %v", got, tt.want)
			}
		})
	}
}
