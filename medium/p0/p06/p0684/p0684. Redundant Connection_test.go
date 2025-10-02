package p0684

import (
	"reflect"
	"testing"
)

func Test_findRedundantConnection(t *testing.T) {
	type args struct {
		edges [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Example 1", args{edges: [][]int{{1, 2}, {1, 3}, {2, 3}}}, []int{2, 3}},
		{"Example 2", args{edges: [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 4}, {1, 5}}}, []int{1, 4}},
		{"TestCase 5", args{edges: [][]int{{1, 2}, {2, 3}, {1, 5}, {3, 4}, {1, 4}}}, []int{1, 4}},
		{"TestCase 12", args{edges: [][]int{{1, 2}, {2, 3}, {2, 4}, {4, 5}, {1, 5}}}, []int{1, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findRedundantConnection(tt.args.edges); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findRedundantConnection() = %v, want %v", got, tt.want)
			}
		})
	}
}
