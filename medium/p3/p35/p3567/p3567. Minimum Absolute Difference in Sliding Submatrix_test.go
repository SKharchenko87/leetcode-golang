package p3567

import (
	"reflect"
	"testing"
)

func Test_minAbsDiff(t *testing.T) {
	type args struct {
		grid [][]int
		k    int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"Example 1", args{grid: [][]int{{1, 8}, {3, -2}}, k: 2}, [][]int{{2}}},
		{"Example 2", args{grid: [][]int{{3, -1}}, k: 1}, [][]int{{0, 0}}},
		{"Example 3", args{grid: [][]int{{1, -2, 3}, {2, 3, 5}}, k: 2}, [][]int{{1, 2}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minAbsDiff(tt.args.grid, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("minAbsDiff() = %v, want %v", got, tt.want)
			}
		})
	}
}
