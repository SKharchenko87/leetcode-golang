package p3619

import "testing"

func Test_countIslands(t *testing.T) {
	type args struct {
		grid [][]int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{grid: [][]int{{0, 2, 1, 0, 0}, {0, 5, 0, 0, 5}, {0, 0, 1, 0, 0}, {0, 1, 4, 7, 0}, {0, 2, 0, 0, 8}}, k: 5}, 2},
		{"Example 2", args{grid: [][]int{{3, 0, 3, 0}, {0, 3, 0, 3}, {3, 0, 3, 0}}, k: 3}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countIslands(tt.args.grid, tt.args.k); got != tt.want {
				t.Errorf("countIslands() = %v, want %v", got, tt.want)
			}
		})
	}
}
