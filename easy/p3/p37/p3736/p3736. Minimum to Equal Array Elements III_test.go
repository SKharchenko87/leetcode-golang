package p3736

import "testing"

func Test_minMoves(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		nums []int
		want int
	}{
		{"Example 1", []int{2, 1, 3}, 3},
		{"Example 2", []int{4, 4, 5}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minMoves(tt.nums)
			if got != tt.want {
				t.Errorf("minMoves() = %v, want %v", got, tt.want)
			}
		})
	}
}
