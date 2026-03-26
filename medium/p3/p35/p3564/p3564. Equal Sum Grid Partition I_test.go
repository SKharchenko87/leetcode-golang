package p3564

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
		{"Example 2", args{grid: [][]int{{1, 3}, {2, 4}}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canPartitionGrid(tt.args.grid); got != tt.want {
				t.Errorf("canPartitionGrid() = %v, want %v", got, tt.want)
			}
		})
	}
}
