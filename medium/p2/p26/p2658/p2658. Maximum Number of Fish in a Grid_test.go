package p2658

import "testing"

func Test_findMaxFish(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{grid: [][]int{{0, 2, 1, 0}, {4, 0, 0, 3}, {1, 0, 0, 4}, {0, 3, 2, 0}}}, 7},
		{"Example 2", args{grid: [][]int{{1, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 1}}}, 1},
		{"TestCase 4", args{grid: [][]int{{0, 4}}}, 4},
		{"TestCase 2863", args{grid: [][]int{{0, 5}, {8, 4}}}, 17},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMaxFish(tt.args.grid); got != tt.want {
				t.Errorf("findMaxFish() = %v, want %v", got, tt.want)
			}
		})
	}
}
