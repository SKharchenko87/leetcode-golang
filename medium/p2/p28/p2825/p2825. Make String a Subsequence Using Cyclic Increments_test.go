package p2825

import "testing"

func Test_canMakeSubsequence(t *testing.T) {
	type args struct {
		str1 string
		str2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Example 1", args{str1: "abc", str2: "ad"}, true},
		{"Example 2", args{str1: "zc", str2: "ad"}, true},
		{"Example 3", args{str1: "ab", str2: "d"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canMakeSubsequence(tt.args.str1, tt.args.str2); got != tt.want {
				t.Errorf("canMakeSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
