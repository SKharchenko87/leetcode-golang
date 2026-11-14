package p3364

import "testing"

func Test_minimumSumSubarray(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		nums []int
		l    int
		r    int
		want int
	}{
		{"Example 1", []int{3, -2, 1, 4}, 2, 3, 1},
		{"Example 2", []int{-2, 2, -3, 1}, 2, 3, -1},
		{"Example 3", []int{1, 2, 3, 4}, 2, 4, 3},
		{"TestCase 426", []int{7, 3}, 2, 2, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minimumSumSubarray(tt.nums, tt.l, tt.r)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("minimumSumSubarray() = %v, want %v", got, tt.want)
			}
		})
	}
}
