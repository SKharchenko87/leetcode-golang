package p2257

import "testing"

func Test_countUnguarded(t *testing.T) {
	type args struct {
		m      int
		n      int
		guards [][]int
		walls  [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{m: 4, n: 6, guards: [][]int{{0, 0}, {1, 1}, {2, 3}}, walls: [][]int{{0, 1}, {2, 2}, {1, 4}}}, 7},
		{"Example 2", args{m: 3, n: 3, guards: [][]int{{1, 1}}, walls: [][]int{{0, 1}, {1, 0}, {2, 1}, {1, 2}}}, 4},
		{"TestCase 21", args{m: 6, n: 10, guards: [][]int{{0, 6}, {2, 2}, {2, 5}, {1, 2}, {4, 9}, {2, 9}, {5, 6}, {4, 6}}, walls: [][]int{{1, 5}}}, 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countUnguarded(tt.args.m, tt.args.n, tt.args.guards, tt.args.walls); got != tt.want {
				t.Errorf("countUnguarded() = %v, want %v", got, tt.want)
			}
		})
	}
}
