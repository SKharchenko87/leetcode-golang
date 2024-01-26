package p1143

import "testing"

func Test_longestCommonSubsequence(t *testing.T) {
	type args struct {
		text1 string
		text2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Case 1", args{"abcde", "ace"}, 3},
		{"Case 2", args{"abc", "abc"}, 3},
		{"Case 3", args{"abc", "def"}, 0},
		{"Case 4", args{"ezupkr", "ubmrapg"}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonSubsequence(tt.args.text1, tt.args.text2); got != tt.want {
				t.Errorf("longestCommonSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
