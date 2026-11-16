package p2243

import "testing"

func Test_digitSum(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		s    string
		k    int
		want string
	}{
		{"Example 1", "11111222223", 3, "135"},
		{"Example 2", "00000000", 3, "000"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := digitSum(tt.s, tt.k)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("digitSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
