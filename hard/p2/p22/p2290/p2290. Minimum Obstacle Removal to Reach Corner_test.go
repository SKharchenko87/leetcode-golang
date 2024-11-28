package p2290

import "testing"

func Test_minimumObstacles(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{grid: [][]int{{0, 1, 1}, {1, 1, 0}, {1, 1, 0}}}, 2},
		{"Example 2", args{grid: [][]int{{0, 1, 0, 0, 0}, {0, 1, 0, 1, 0}, {0, 0, 0, 1, 0}}}, 0},
		{"TestCase 2", args{grid: [][]int{{0, 1, 1, 1, 0}}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumObstacles(tt.args.grid); got != tt.want {
				t.Errorf("minimumObstacles() = %v, want %v", got, tt.want)
			}
		})
	}
}
