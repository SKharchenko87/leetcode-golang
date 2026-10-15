package p0856

import "testing"

func Test_scoreParentheses(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{s: "()"}, 1},
		{"Example 2", args{s: "(())"}, 2},
		{"Example 3", args{s: "()()"}, 2},
		{"My 1", args{s: "(()())"}, 4},
		{"My 1", args{s: "((()()))"}, 8},
		{"My 1", args{s: "((()())())"}, 10},
		{"TestCase 57", args{s: "(()(()))"}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := scoreOfParentheses(tt.args.s); got != tt.want {
				t.Errorf("scoreOfParentheses() = %v, want %v", got, tt.want)
			}
		})
	}
}
