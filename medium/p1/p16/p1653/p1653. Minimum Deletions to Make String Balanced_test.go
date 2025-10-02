package p1653

import "testing"

func Test_minimumDeletions(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{s: "aababbab"}, 2},
		{"Example 2", args{s: "bbaaaaabb"}, 2},
		{"Example 2", args{s: "bbaaaaabba"}, 3},
		{"Testcase 154", args{s: "bbbbbbbbbbbbbb"}, 0},
		{"Testcase 155", args{s: "bbbbbbbaabbbbbaaabbbabbbbaabbbbbbaabbaaabaabbbaaaabaaababbbabbabbaaaabbbabbbbbaabbababbbaaaaaababaaababaabbabbbaaaabbbbbabbabaaaabbbaba"}, 60},
		{"My 1", args{s: "bbabaaabbb"}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumDeletions(tt.args.s); got != tt.want {
				t.Errorf("minimumDeletions() = %v, want %v", got, tt.want)
			}
		})
	}
}
