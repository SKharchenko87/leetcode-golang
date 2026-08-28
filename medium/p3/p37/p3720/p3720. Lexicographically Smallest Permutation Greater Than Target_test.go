package p3720

import "testing"

func Test_lexGreaterPermutation(t *testing.T) {
	type args struct {
		s      string
		target string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Example 1", args{s: "abc", target: "bba"}, "bca"},
		{"Example 2", args{s: "leet", target: "code"}, "eelt"},
		{"Example 3", args{s: "baba", target: "bbaa"}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lexGreaterPermutation(tt.args.s, tt.args.target); got != tt.want {
				t.Errorf("lexGreaterPermutation() = %v, want %v", got, tt.want)
			}
		})
	}
}
