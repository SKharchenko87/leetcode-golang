package p3693

import "testing"

func Test_climbStairs(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		n     int
		costs []int
		want  int
	}{
		{"Example 1", 4, []int{1, 2, 3, 4}, 13},
		{"Example 2", 4, []int{5, 1, 6, 2}, 11},
		{"Example 3", 3, []int{9, 8, 3}, 12}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := climbStairs(tt.n, tt.costs)
			if got != tt.want {
				t.Errorf("climbStairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
