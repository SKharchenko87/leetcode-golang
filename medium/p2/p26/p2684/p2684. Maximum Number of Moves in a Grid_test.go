package p2684

import "testing"

func Test_maxMoves(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{grid: [][]int{{2, 4, 3, 5}, {5, 4, 9, 3}, {3, 4, 2, 11}, {10, 9, 13, 15}}}, 3},
		{"Example 2", args{grid: [][]int{{3, 2, 4}, {2, 1, 9}, {1, 1, 7}}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxMoves(tt.args.grid); got != tt.want {
				t.Errorf("maxMoves() = %v, want %v", got, tt.want)
			}
		})
	}
}
