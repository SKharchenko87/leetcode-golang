package p2872

import "testing"

func Test_maxKDivisibleComponents(t *testing.T) {
	type args struct {
		n      int
		edges  [][]int
		values []int
		k      int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{n: 5, edges: [][]int{{0, 2}, {1, 2}, {1, 3}, {2, 4}}, values: []int{1, 8, 1, 4, 4}, k: 6}, 2},
		{"Example 2", args{n: 7, edges: [][]int{{0, 1}, {0, 2}, {1, 3}, {1, 4}, {2, 5}, {2, 6}}, values: []int{3, 0, 6, 1, 5, 2, 1}, k: 3}, 3},
		{"TestCase 57", args{n: 2, edges: [][]int{{1, 0}}, values: []int{0, 0}, k: 100000000}, 2},
		{"TestCase 734", args{n: 1, edges: [][]int{}, values: []int{0}, k: 1}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxKDivisibleComponents(tt.args.n, tt.args.edges, tt.args.values, tt.args.k); got != tt.want {
				t.Errorf("maxKDivisibleComponents() = %v, want %v", got, tt.want)
			}
		})
	}
}
