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

func Test_countSubmatrices(t *testing.T) {
	type args struct {
		grid [][]int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{[][]int{{7, 6, 3}, {6, 6, 1}}, 18}, 4},
		{"Example 2", args{[][]int{{7, 2, 9}, {1, 5, 0}, {2, 6, 6}}, 20}, 6},
		{"Example 2", args{[][]int{{7, 2, 9}, {1, 5, 0}, {2, 6, 6}}, 100}, 9},
		{"Example 3", args{[][]int{{9, 5}, {10, 8}, {7, 6}}, 3}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countSubmatrices(tt.args.grid, tt.args.k); got != tt.want {
				t.Errorf("countSubmatrices() = %v, want %v", got, tt.want)
			}
		})
	}
}
