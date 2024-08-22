package p0664

import "testing"

func Test_strangePrinter(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{"aaabbb"}, 2},
		{"Example 2", args{"aba"}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strangePrinter(tt.args.s); got != tt.want {
				t.Errorf("strangePrinter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_removeDuplicates(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"My 0", args{"a"}, "a"},
		{"My 1", args{"leetcode"}, "letcode"},
		{"My 2", args{"aabcdaa"}, "abcda"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates(tt.args.s); got != tt.want {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
