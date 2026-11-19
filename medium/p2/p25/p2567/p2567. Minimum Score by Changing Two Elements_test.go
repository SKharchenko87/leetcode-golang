package p2567

import "testing"

func Test_minimizeSum(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		nums []int
		want int
	}{
		{"Example 1", []int{1, 4, 7, 8, 5}, 3},
		{"Example 2", []int{1, 4, 3}, 0},
		{"TestCase 53", []int{21, 13, 21, 72, 35, 52, 74}, 39},
		{"TestCase 54", []int{40, 60, 23, 99, 83, 5, 21, 76, 34, 99}, 76},
		{"TestCase 58", []int{23, 70, 18, 64, 12, 21, 21, 78, 36, 58}, 52},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minimizeSum(tt.nums)
			if got != tt.want {
				t.Errorf("minimizeSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
