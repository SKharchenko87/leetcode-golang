package p2

import "testing"

func Test_abs(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Case 1", args{-1}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := abs(tt.args.x); got != tt.want {
				t.Errorf("abs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lastNonEmptyString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Example 1", args{"aabcbbca"}, "ba"},
		{"Example 2", args{"abcd"}, "abcd"},
		{"Case 1", args{"aaaabcd"}, "a"},
		{"Case 1", args{"bbbaaa"}, "ba"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lastNonEmptyString(tt.args.s); got != tt.want {
				t.Errorf("lastNonEmptyString() = %v, want %v", got, tt.want)
			}
		})
	}
}
