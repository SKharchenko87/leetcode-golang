package p2461

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
		{"Example 1", args{nums: []int{1, 5, 4, 2, 9, 9, 9}, k: 3}, 15},
		{"Example 2", args{nums: []int{4, 4, 4}, k: 3}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumSubarraySum(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("maximumSubarraySum() = %v, want %v", got, tt.want)
			}
		})
	}
}
