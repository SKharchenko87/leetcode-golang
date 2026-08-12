package p2958

import "testing"

func Test_maxSubarrayLength(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{nums: []int{1, 2, 3, 1, 2, 3, 1, 2}, k: 2}, 6},
		{"Example 2", args{nums: []int{1, 2, 1, 2, 1, 2, 1, 2}, k: 1}, 2},
		{"Example 3", args{nums: []int{5, 5, 5, 5, 5, 5, 5}, k: 4}, 4},
		{"Example 3", args{nums: []int{1, 4, 4, 3}, k: 1}, 2},
		{"Example 3", args{nums: []int{3, 1, 1}, k: 1}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSubarrayLength(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("maxSubarrayLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
