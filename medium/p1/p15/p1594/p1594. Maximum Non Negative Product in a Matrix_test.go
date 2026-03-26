package p1594

import "testing"

func Test_maxProductPath(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{grid: [][]int{{-1, -2, -3}, {-2, -3, -3}, {-3, -3, -2}}}, -1},
		{"Example 2", args{grid: [][]int{{1, -2, 1}, {1, -2, 1}, {3, -4, 1}}}, 8},
		{"Example 3", args{grid: [][]int{{1, 3}, {0, -4}}}, 0},
		{"TestCase 135", args{grid: [][]int{{-1, 3, 0}, {3, -2, 3}, {-1, 1, -4}}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProductPath(tt.args.grid); got != tt.want {
				t.Errorf("maxProductPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
