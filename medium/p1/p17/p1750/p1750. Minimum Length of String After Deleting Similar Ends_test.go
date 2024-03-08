package p1750

import "testing"

func Test_minimumLength(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{"ca"}, 2},
		{"Example 2", args{"cabaabac"}, 0},
		{"Example 3", args{"aabccabba"}, 3},
		{"Example 3", args{"bbbbbbbbbbbbbbbbbbbbbbbbbbbabbbbbbbbbbbbbbbccbcbcbccbbabbb"}, 1},
		{"Example 3", args{"bbbbbbbbbbbbbbbbbbb"}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumLength(tt.args.s); got != tt.want {
				t.Errorf("minimumLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
