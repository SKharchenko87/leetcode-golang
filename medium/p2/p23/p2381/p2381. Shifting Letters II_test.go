package p2381

import "testing"

func Test_shiftingLetters(t *testing.T) {
	type args struct {
		s      string
		shifts [][]int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Example 1", args{s: "abc", shifts: [][]int{{0, 1, 0}, {1, 2, 1}, {0, 2, 1}}}, "ace"},
		{"Example 2", args{s: "dztz", shifts: [][]int{{0, 0, 0}, {1, 1, 1}}}, "catz"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shiftingLetters(tt.args.s, tt.args.shifts); got != tt.want {
				t.Errorf("shiftingLetters() = %v, want %v", got, tt.want)
			}
		})
	}
}
