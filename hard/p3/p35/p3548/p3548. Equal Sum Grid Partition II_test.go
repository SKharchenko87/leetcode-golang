package p3548

import "testing"

func Test_canPartitionGrid(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Example 1", args{grid: [][]int{{1, 4}, {2, 3}}}, true},
		{"Example 2", args{grid: [][]int{{1, 2}, {3, 4}}}, true},
		{"Example 3", args{grid: [][]int{{1, 2, 4}, {2, 3, 5}}}, false},
		{"Example 4", args{grid: [][]int{{4, 1, 8}, {3, 2, 6}}}, false},
		{"Example 5", args{grid: [][]int{{25372}, {100000}, {100000}}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canPartitionGrid(tt.args.grid); got != tt.want {
				t.Errorf("canPartitionGrid() = %v, want %v", got, tt.want)
			}
		})
	}
}
