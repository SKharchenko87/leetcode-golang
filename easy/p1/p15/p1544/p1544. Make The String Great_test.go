package p1544

import "testing"

func Test_makeGood(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Example 1", args{"leEeetcode"}, "leetcode"},
		{"Example 2", args{"abBAcC"}, ""},
		{"Example 3", args{"s"}, "s"},
		{"Example 74", args{"Pp"}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeGood(tt.args.s); got != tt.want {
				t.Errorf("makeGood() = %v, want %v", got, tt.want)
			}
		})
	}
}
