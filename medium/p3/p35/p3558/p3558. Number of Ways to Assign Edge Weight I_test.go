package p3558

import "testing"

func Test_assignEdgeWeight(t *testing.T) {
	type args struct {
		edges [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{edges: [][]int{{1, 2}}}, 1},
		{"Example 2", args{edges: [][]int{{1, 2}, {1, 3}, {3, 4}, {3, 5}}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := assignEdgeWeights(tt.args.edges); got != tt.want {
				t.Errorf("assignEdgeWeights() = %v, want %v", got, tt.want)
			}
		})
	}
}
