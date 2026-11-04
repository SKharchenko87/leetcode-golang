package p3505

import "testing"

func Test_minOperations(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		nums []int
		x    int
		k    int
		want int64
	}{
		{"Example 1", []int{5, -2, 1, 3, 7, 3, 6, 4, -1}, 3, 2, 8},
		{"Example 2", []int{9, -2, -2, -2, 1, 5}, 2, 2, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minOperations(tt.nums, tt.x, tt.k)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("minOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
