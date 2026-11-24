package p3411

import "testing"

func Test_maxLength(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		nums []int
		want int
	}{
		{"Example 1", []int{1, 2, 1, 2, 1, 1, 1}, 5},
		{"Example 2", []int{2, 3, 4, 5, 6}, 3},
		{"Example 3", []int{1, 2, 3, 1, 4, 5, 1}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxLength(tt.nums)
			if got != tt.want {
				t.Errorf("maxLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
