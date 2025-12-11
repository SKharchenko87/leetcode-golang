package p3531

import "testing"

func Test_countCoveredBuildings(t *testing.T) {
	type args struct {
		n         int
		buildings [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{n: 3, buildings: [][]int{{1, 2}, {2, 2}, {3, 2}, {2, 1}, {2, 3}}}, 1},
		{"Example 2", args{n: 3, buildings: [][]int{{1, 1}, {1, 2}, {2, 1}, {2, 2}}}, 0},
		{"Example 3", args{n: 5, buildings: [][]int{{1, 3}, {3, 2}, {3, 3}, {3, 5}, {5, 3}}}, 1},
		{"TestCase 341", args{n: 4, buildings: [][]int{{2, 4}, {1, 2}, {3, 1}, {1, 4}, {2, 3}, {3, 3}, {2, 2}, {1, 3}}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countCoveredBuildings(tt.args.n, tt.args.buildings); got != tt.want {
				t.Errorf("countCoveredBuildings() = %v, want %v", got, tt.want)
			}
		})
	}
}
