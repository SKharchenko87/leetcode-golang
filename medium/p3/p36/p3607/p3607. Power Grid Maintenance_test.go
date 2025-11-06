package p3607

import (
	"reflect"
	"testing"
)

func Test_processQueries(t *testing.T) {
	type args struct {
		c           int
		connections [][]int
		queries     [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Example 1", args{c: 5, connections: [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}}, queries: [][]int{{1, 3}, {2, 1}, {1, 1}, {2, 2}, {1, 2}}}, []int{3, 2, 3}},
		{"Example 2", args{c: 3, connections: [][]int{}, queries: [][]int{{1, 1}, {2, 1}, {1, 1}}}, []int{1, -1}},
		{"TestCase 464", args{c: 2, connections: [][]int{{1, 2}}, queries: [][]int{{1, 2}, {2, 1}, {1, 2}, {1, 1}, {2, 2}, {1, 2}, {2, 2}, {1, 1}, {1, 1}, {2, 1}, {1, 2}, {1, 2}, {1, 1}, {2, 1}, {1, 1}, {1, 1}, {1, 1}, {2, 1}, {2, 1}, {2, 1}, {1, 2}, {2, 1}, {2, 2}, {1, 2}, {1, 1}, {2, 2}, {2, 1}, {2, 2}}}, []int{2, 2, 2, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1}},
		{"TestCase 575", args{c: 4, connections: [][]int{{2, 3}, {1, 3}, {4, 1}, {3, 4}}, queries: [][]int{{1, 2}, {2, 4}, {2, 1}, {1, 4}, {2, 1}, {1, 1}, {2, 2}, {1, 4}, {2, 1}, {2, 2}, {2, 3}, {2, 4}, {2, 1}, {1, 1}, {2, 3}, {2, 2}, {2, 3}, {1, 4}, {2, 4}}}, []int{2, 2, 2, 3, -1, -1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := processQueries(tt.args.c, tt.args.connections, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("processQueries() = %v, want %v", got, tt.want)
			}
		})
	}
}
