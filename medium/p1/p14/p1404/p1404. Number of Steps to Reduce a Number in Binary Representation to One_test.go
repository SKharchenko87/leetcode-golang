package p1404

import "testing"

func Test_numSteps(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{"1101"}, 6},
		{"Example 2", args{"10"}, 1},
		{"Example 3", args{"1"}, 0},
		{"My 1", args{"1000111"}, 11},
		{"My 2", args{"11"}, 3},
		{"My 3", args{"100"}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numSteps(tt.args.s); got != tt.want {
				t.Errorf("numSteps() = %v, want %v", got, tt.want)
			}
		})
	}
}
