package p3459

import "testing"

func Test_lenOfVDiagonal(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{grid: [][]int{{2, 2, 1, 2, 2}, {2, 0, 2, 2, 0}, {2, 0, 1, 1, 0}, {1, 0, 2, 2, 2}, {2, 0, 0, 2, 2}}}, 5},
		{"Example 2", args{grid: [][]int{{2, 2, 2, 2, 2}, {2, 0, 2, 2, 0}, {2, 0, 1, 1, 0}, {1, 0, 2, 2, 2}, {2, 0, 0, 2, 2}}}, 4},
		{"Example 3", args{grid: [][]int{{1, 2, 2, 2, 2}, {2, 2, 2, 2, 0}, {2, 0, 0, 0, 0}, {0, 0, 2, 2, 2}, {2, 0, 0, 2, 0}}}, 5},
		{"Example 4", args{grid: [][]int{{1}}}, 1},
		{"TestCase 54", args{grid: [][]int{{1, 2, 2}, {1, 0, 2}}}, 2},
		{"TestCase 53", args{grid: [][]int{{1, 1, 1, 2, 0, 0}, {0, 0, 0, 0, 1, 2}}}, 2},
		{"TestCase 515", args{grid: [][]int{{0, 2, 2, 2, 0, 2, 2, 2, 0, 2, 1, 2, 0, 1, 0, 0}, {1, 2, 0, 1, 2, 2, 0, 2, 0, 2, 0, 2, 2, 2, 2, 0}, {0, 2, 2, 0, 0, 2, 0, 0, 2, 2, 0, 0, 0, 1, 0, 1}, {0, 1, 0, 2, 0, 1, 1, 2, 0, 2, 0, 0, 0, 1, 0, 2}, {0, 2, 2, 2, 0, 2, 0, 0, 0, 2, 0, 2, 0, 2, 0, 0}}}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lenOfVDiagonal(tt.args.grid); got != tt.want {
				t.Errorf("lenOfVDiagonal() = %v, want %v", got, tt.want)
			}
		})
	}
}
