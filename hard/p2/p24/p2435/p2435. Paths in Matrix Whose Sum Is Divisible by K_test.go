package p2435

import "testing"

func Test_numberOfPaths(t *testing.T) {
	type args struct {
		grid [][]int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{grid: [][]int{{5, 2, 4}, {3, 0, 5}, {0, 7, 2}}, k: 3}, 2},
		{"Example 2", args{grid: [][]int{{0, 0}}, k: 5}, 1},
		{"Example 3", args{grid: [][]int{{7, 3, 4, 9}, {2, 3, 6, 2}, {2, 3, 7, 0}}, k: 1}, 10},
		{"TestCase 27", args{grid: [][]int{{1, 5, 3, 7, 3, 2, 3, 5}}, k: 29}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfPaths(tt.args.grid, tt.args.k); got != tt.want {
				t.Errorf("numberOfPaths() = %v, want %v", got, tt.want)
			}
		})
	}
}
