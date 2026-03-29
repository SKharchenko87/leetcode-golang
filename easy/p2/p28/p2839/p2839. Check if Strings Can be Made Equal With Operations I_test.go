package p2839

import "testing"

func Test_canBeEqual(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Example 1", args{s1: "abcd", s2: "cdab"}, true},
		{"Example 2", args{s1: "abcd", s2: "dacb"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canBeEqual(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("canBeEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}
