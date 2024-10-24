package p2938

import "testing"

func Test_minimumSteps(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"Example 1", args{s: "101"}, 1},
		{"Example 2", args{s: "100"}, 2},
		{"Example 3", args{s: "0111"}, 0},
		{"My 1", args{s: "11000"}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumSteps(tt.args.s); got != tt.want {
				t.Errorf("minimumSteps() = %v, want %v", got, tt.want)
			}
		})
	}
}
