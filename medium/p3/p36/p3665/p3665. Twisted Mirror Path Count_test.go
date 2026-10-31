package p3665

import "testing"

func Test_uniquePaths(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{grid: [][]int{{0, 1, 0}, {0, 0, 1}, {1, 0, 0}}}, 5},
		{"Example 2", args{grid: [][]int{{0, 0}, {0, 0}}}, 2},
		{"Example 3", args{grid: [][]int{{0, 1, 1}, {1, 1, 0}}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniquePaths(tt.args.grid); got != tt.want {
				t.Errorf("uniquePaths() = %v, want %v", got, tt.want)
			}
		})
	}
}
