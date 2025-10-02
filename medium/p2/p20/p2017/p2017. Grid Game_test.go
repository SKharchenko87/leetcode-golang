package p2017

import "testing"

func Test_gridGame(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"Example 1", args{grid: [][]int{{2, 5, 4}, {1, 5, 1}}}, 4},
		{"Example 2", args{grid: [][]int{{3, 3, 1}, {8, 5, 2}}}, 4},
		{"Example 3", args{grid: [][]int{{1, 3, 1, 15}, {1, 3, 3, 1}}}, 7},
		{"TestCase 9", args{grid: [][]int{{20, 3, 20, 17, 2, 12, 15, 17, 4, 15}, {20, 10, 13, 14, 15, 5, 2, 3, 14, 3}}}, 63},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gridGame(tt.args.grid); got != tt.want {
				t.Errorf("gridGame() = %v, want %v", got, tt.want)
			}
		})
	}
}
