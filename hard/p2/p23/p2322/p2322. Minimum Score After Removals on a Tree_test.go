package p2322

import "testing"

func Test_minimumScore(t *testing.T) {
	type args struct {
		nums  []int
		edges [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{nums: []int{1, 5, 5, 4, 11}, edges: [][]int{{0, 1}, {1, 2}, {1, 3}, {3, 4}}}, 9},
		{"Example 2", args{nums: []int{5, 5, 2, 4, 4, 2}, edges: [][]int{{0, 1}, {1, 2}, {5, 2}, {4, 3}, {1, 3}}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumScore(tt.args.nums, tt.args.edges); got != tt.want {
				t.Errorf("minimumScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
