package p3243

import (
	"reflect"
	"testing"
)

func Test_shortestDistanceAfterQueries(t *testing.T) {
	type args struct {
		n       int
		queries [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Example 1", args{n: 5, queries: [][]int{{2, 4}, {0, 2}, {0, 4}}}, []int{3, 2, 1}},
		{"Example 2", args{n: 4, queries: [][]int{{0, 3}, {0, 2}}}, []int{1, 1}},
		{"TestCase 308", args{n: 6, queries: [][]int{{1, 4}, {2, 4}}}, []int{3, 3}},
		{"TestCase 373", args{n: 8, queries: [][]int{{5, 7}, {0, 6}}}, []int{6, 2}},
		{"TestCase 877", args{n: 13, queries: [][]int{{9, 11}, {2, 6}, {6, 9}}}, []int{11, 8, 6}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shortestDistanceAfterQueries(tt.args.n, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("shortestDistanceAfterQueries() = %v, want %v", got, tt.want)
			}
		})
	}
}
