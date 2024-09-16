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

func Test_findSafeWalk(t *testing.T) {
	type args struct {
		grid   [][]int
		health int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Example 1", args{grid: [][]int{{0, 1, 0, 0, 0}, {0, 1, 0, 1, 0}, {0, 0, 0, 1, 0}}, health: 1}, true},
		{"Example 2", args{grid: [][]int{{0, 1, 1, 0, 0, 0}, {1, 0, 1, 0, 0, 0}, {0, 1, 1, 1, 0, 1}, {0, 0, 1, 0, 1, 0}}, health: 3}, false},
		{"Example 3", args{grid: [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}, health: 5}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findSafeWalk(tt.args.grid, tt.args.health); got != tt.want {
				t.Errorf("findSafeWalk() = %v, want %v", got, tt.want)
			}
		})
	}
}
