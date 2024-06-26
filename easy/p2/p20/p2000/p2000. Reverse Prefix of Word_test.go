package p2000

import "testing"

func Test_reversePrefix(t *testing.T) {
	type args struct {
		word string
		ch   byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Example 1", args{"abcdefd", 'd'}, "dcbaefd"},
		{"Example 2", args{"xyxzxe", 'z'}, "zxyxxe"},
		{"Example 3", args{"abcd", 'z'}, "abcd"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reversePrefix(tt.args.word, tt.args.ch); got != tt.want {
				t.Errorf("reversePrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
