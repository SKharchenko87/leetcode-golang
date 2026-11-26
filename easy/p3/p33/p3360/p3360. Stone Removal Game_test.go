package p3360

import "testing"

func Test_canAliceWin(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		n    int
		want bool
	}{
		{"Example 1", 12, true},
		{"Example 2", 1, false},
		{"TestCase 46", 10, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := canAliceWin(tt.n)
			if got != tt.want {
				t.Errorf("canAliceWin() = %v, want %v", got, tt.want)
			}
		})
	}
}
