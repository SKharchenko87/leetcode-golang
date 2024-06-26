package p0834

import (
	"reflect"
	"testing"
)

func Test_sumOfDistancesInTree(t *testing.T) {
	type args struct {
		n     int
		edges [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Example 1", args{6, [][]int{{0, 1}, {0, 2}, {2, 3}, {2, 4}, {2, 5}}}, []int{8, 12, 6, 10, 10, 10}},
		{"Example 2", args{1, [][]int{}}, []int{0}},
		{"Example 3", args{2, [][]int{{1, 0}}}, []int{1, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumOfDistancesInTree(tt.args.n, tt.args.edges); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sumOfDistancesInTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
