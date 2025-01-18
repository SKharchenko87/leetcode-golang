package p1368

import "testing"

func Test_minCost(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{grid: [][]int{{1, 1, 1, 1}, {2, 2, 2, 2}, {1, 1, 1, 1}, {2, 2, 2, 2}}}, 3},
		{"Example 2", args{grid: [][]int{{1, 1, 3}, {3, 2, 2}, {1, 1, 4}}}, 0},
		{"Example 3", args{grid: [][]int{{1, 2}, {4, 3}}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCost(tt.args.grid); got != tt.want {
				t.Errorf("minCost() = %v, want %v", got, tt.want)
			}
		})
	}
}
