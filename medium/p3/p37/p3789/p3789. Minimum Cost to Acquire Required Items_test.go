package p3789

import "testing"

func Test_minimumCost(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		cost1    int
		cost2    int
		costBoth int
		need1    int
		need2    int
		want     int64
	}{
		{"Example 1", 3, 2, 1, 3, 2, 3},
		{"Example 2", 5, 4, 15, 2, 3, 22},
		{"Example 3", 5, 4, 15, 0, 0, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minimumCost(tt.cost1, tt.cost2, tt.costBoth, tt.need1, tt.need2)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("minimumCost() = %v, want %v", got, tt.want)
			}
		})
	}
}
