package p2699

import (
	"reflect"
	"testing"
)

func Test_modifiedGraphEdges(t *testing.T) {
	type args struct {
		n           int
		edges       [][]int
		source      int
		destination int
		target      int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"Example 1", args{n: 5, edges: [][]int{{4, 1, -1}, {2, 0, -1}, {0, 3, -1}, {4, 3, -1}}, source: 0, destination: 1, target: 5}, [][]int{{4, 1, 1}, {2, 0, 1}, {0, 3, 3}, {4, 3, 1}}},
		{"Example 2", args{n: 3, edges: [][]int{{0, 1, -1}, {0, 2, 5}}, source: 0, destination: 2, target: 6}, [][]int{}},
		{"Example 3", args{n: 4, edges: [][]int{{1, 0, 4}, {1, 2, 3}, {2, 3, 5}, {0, 3, -1}}, source: 0, destination: 2, target: 6}, [][]int{{1, 0, 4}, {1, 2, 3}, {2, 3, 5}, {0, 3, 1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := modifiedGraphEdges(tt.args.n, tt.args.edges, tt.args.source, tt.args.destination, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("modifiedGraphEdges() = %v, want %v", got, tt.want)
			}
		})
	}
}
