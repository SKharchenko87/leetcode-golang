package p3

import "testing"

func Test_maximumSubarraySum(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"Case 1", args{[]int{1, 2, 3, 4, 5, 6}, 1}, 11},
		{"Case 2", args{[]int{-1, 3, 2, 4, 5}, 3}, 11},
		{"Case 3", args{[]int{-1, -2, -3, -4}, 2}, -6},
		{"Case 3", args{[]int{3, 3, 2}, 1}, 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumSubarraySum(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("maximumSubarraySum() = %v, want %v", got, tt.want)
			}
		})
	}
}
