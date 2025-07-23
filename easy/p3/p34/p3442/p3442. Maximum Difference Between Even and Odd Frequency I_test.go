package p3442

import "testing"

func Test_maxDifference(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{s: "aaaaabbc"}, 3},
		{"Example 2", args{s: "abcabcab"}, 1},
		{"Testcase 497", args{s: "ililm"}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDifference(tt.args.s); got != tt.want {
				t.Errorf("maxDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
