package p2924

import "testing"

func Test_findChampion(t *testing.T) {
	type args struct {
		n     int
		edges [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{n: 3, edges: [][]int{{0, 1}, {1, 2}}}, 0},
		{"Example 2", args{n: 4, edges: [][]int{{0, 2}, {1, 3}, {1, 2}}}, -1},
		{"Example 486", args{n: 3, edges: [][]int{{0, 1}, {1, 2}, {0, 2}}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findChampion(tt.args.n, tt.args.edges); got != tt.want {
				t.Errorf("findChampion() = %v, want %v", got, tt.want)
			}
		})
	}
}
