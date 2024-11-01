package p1957

import "testing"

func Test_makeFancyString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Example 1", args{s: "leeetcode"}, "leetcode"},
		{"Example 2", args{s: "aaabaaaa"}, "aabaa"},
		{"Example 3", args{s: "aab"}, "aab"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeFancyString(tt.args.s); got != tt.want {
				t.Errorf("makeFancyString() = %v, want %v", got, tt.want)
			}
		})
	}
}
