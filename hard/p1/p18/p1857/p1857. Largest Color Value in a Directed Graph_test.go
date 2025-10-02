package p1857

import "testing"

func Test_largestPathValue(t *testing.T) {
	type args struct {
		colors string
		edges  [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{colors: "abaca", edges: [][]int{{0, 1}, {0, 2}, {2, 3}, {3, 4}}}, 3},
		{"Example 2", args{colors: "a", edges: [][]int{{0, 0}}}, -1},
		{"TestCase 19", args{colors: "hhqhuqhqff", edges: [][]int{{0, 1}, {0, 2}, {2, 3}, {3, 4}, {3, 5}, {5, 6}, {2, 7}, {6, 7}, {7, 8}, {3, 8}, {5, 8}, {8, 9}, {3, 9}, {6, 9}}}, 3},
		{"TestCase 39", args{colors: "bbbhb", edges: [][]int{{0, 2}, {3, 0}, {1, 3}, {4, 1}}}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestPathValue(tt.args.colors, tt.args.edges); got != tt.want {
				t.Errorf("largestPathValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
