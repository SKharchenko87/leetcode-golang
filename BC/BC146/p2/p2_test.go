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

func Test_countPathsWithXorValue(t *testing.T) {
	type args struct {
		grid [][]int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{grid: [][]int{{2, 1, 5}, {7, 10, 0}, {12, 6, 4}}, k: 11}, 3},
		{"Example 2", args{grid: [][]int{{1, 3, 3, 3}, {0, 3, 3, 2}, {3, 0, 1, 1}}, k: 2}, 5},
		{"Example 3", args{grid: [][]int{{1, 1, 1, 2}, {3, 0, 3, 2}, {3, 0, 2, 2}}, k: 10}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countPathsWithXorValue(tt.args.grid, tt.args.k); got != tt.want {
				t.Errorf("countPathsWithXorValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
