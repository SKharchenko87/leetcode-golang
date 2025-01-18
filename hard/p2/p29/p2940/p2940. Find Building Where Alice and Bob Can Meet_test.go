package p2940

import (
	"reflect"
	"testing"
)

func Test_leftmostBuildingQueries(t *testing.T) {
	type args struct {
		heights []int
		queries [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Example 1", args{heights: []int{6, 4, 8, 5, 2, 7}, queries: [][]int{{0, 1}, {0, 3}, {2, 4}, {3, 4}, {2, 2}}}, []int{2, 5, -1, 5, 2}},
		{"Example 2", args{heights: []int{5, 3, 8, 2, 6, 1, 4, 6}, queries: [][]int{{0, 7}, {3, 5}, {5, 2}, {3, 0}, {1, 6}}}, []int{7, 6, -1, 4, 6}},
		{"Example 2", args{heights: []int{1, 2}, queries: [][]int{{0, 0}, {0, 1}, {1, 0}, {1, 1}}}, []int{0, 1, 1, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := leftmostBuildingQueries(tt.args.heights, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("leftmostBuildingQueries() = %v, want %v", got, tt.want)
			}
		})
	}
}
