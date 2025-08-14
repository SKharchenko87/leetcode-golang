package p3203

import "testing"

func Test_minimumDiameterAfterMerge(t *testing.T) {
	type args struct {
		edges1 [][]int
		edges2 [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{edges1: [][]int{{0, 1}, {0, 2}, {0, 3}}, edges2: [][]int{{0, 1}}}, 3},
		{"Example 2", args{edges1: [][]int{{0, 1}, {0, 2}, {0, 3}, {2, 4}, {2, 5}, {3, 6}, {2, 7}}, edges2: [][]int{{0, 1}, {0, 2}, {0, 3}, {2, 4}, {2, 5}, {3, 6}, {2, 7}}}, 5},
		{"TestCase 216", args{edges1: [][]int{{0, 1}}, edges2: [][]int{}}, 2},
		{"TestCase XXX", args{edges1: [][]int{{0, 1}, {2, 0}, {3, 2}, {3, 6}, {8, 7}, {4, 8}, {5, 4}, {3, 5}, {3, 9}}, edges2: [][]int{{0, 1}, {0, 2}, {0, 3}}}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumDiameterAfterMerge(tt.args.edges1, tt.args.edges2); got != tt.want {
				t.Errorf("minimumDiameterAfterMerge() = %v, want %v", got, tt.want)
			}
		})
	}
}
