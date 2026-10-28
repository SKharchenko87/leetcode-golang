package p3882

import "testing"

func Test_minCost(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{grid: [][]int{{1, 2}, {3, 4}}}, 6},
		{"Example 2", args{grid: [][]int{{6, 7}, {5, 8}}}, 9},
		{"Example 3", args{grid: [][]int{{2, 7, 5}}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCost(tt.args.grid); got != tt.want {
				t.Errorf("minCost() = %v, want %v", got, tt.want)
			}
		})
	}
}
