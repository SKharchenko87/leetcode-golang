package p1025

import "testing"

func Test_divisorGame(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Example 1", args{n: 2}, true},
		{"Example 2", args{n: 3}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := divisorGame(tt.args.n); got != tt.want {
				t.Errorf("divisorGame() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_divisorGame0(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Example 1", args{n: 2}, true},
		{"Example 2", args{n: 3}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := divisorGame0(tt.args.n); got != tt.want {
				t.Errorf("divisorGame0() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_divisorGame1(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Example 1", args{n: 2}, true},
		{"Example 2", args{n: 3}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := divisorGame1(tt.args.n); got != tt.want {
				t.Errorf("divisorGame1() = %v, want %v", got, tt.want)
			}
		})
	}
}
