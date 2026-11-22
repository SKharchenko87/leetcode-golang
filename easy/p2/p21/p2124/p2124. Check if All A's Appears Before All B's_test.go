package p2124

import "testing"

func Test_checkString(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		s    string
		want bool
	}{
		{"Example 1", "aaabbb", true},
		{"Example 2", "abab", false},
		{"Example 3", "bbb", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := checkString(tt.s)
			if got != tt.want {
				t.Errorf("checkString() = %v, want %v", got, tt.want)
			}
		})
	}
}
