package p4

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

func Test_minStartingIndex(t *testing.T) {
	type args struct {
		s       string
		pattern string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"T1", args{"abcdefg", "bcdffg"}, 1},
		{"T2", args{"ababbababa", "bacaba"}, 4},
		{"T3", args{"abcd", "dba"}, -1},
		{"T4", args{"dde", "d"}, 0},
		{"T677", args{"ccbb", "bc"}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minStartingIndex(tt.args.s, tt.args.pattern); got != tt.want {
				t.Errorf("minStartingIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
