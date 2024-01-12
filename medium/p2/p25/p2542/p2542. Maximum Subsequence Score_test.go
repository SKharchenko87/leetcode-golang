package p2542

import "testing"

func Test_maxScore(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
		k     int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"Case 1", args{[]int{1, 3, 3, 2}, []int{2, 1, 3, 4}, 3}, 12},
		{"Case 11", args{[]int{2, 1, 14, 12}, []int{11, 7, 13, 6}, 3}, 168},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxScore(tt.args.nums1, tt.args.nums2, tt.args.k); got != tt.want {
				t.Errorf("maxScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
