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

func Test_minimumTimeToInitialState(t *testing.T) {
	type args struct {
		word string
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{

		{"1", args{"abacabaabac", 4}, 3},
		{"1", args{"abacaba", 4}, 1},
		{"1", args{"abacaba", 3}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumTimeToInitialState(tt.args.word, tt.args.k); got != tt.want {
				t.Errorf("minimumTimeToInitialState() = %v, want %v", got, tt.want)
			}
		})
	}
}
