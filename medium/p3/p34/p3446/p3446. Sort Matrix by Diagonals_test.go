package p3446

import (
	"reflect"
	"testing"
)

func Test_sortMatrix(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"Example 1", args{grid: [][]int{{1, 7, 3}, {9, 8, 2}, {4, 5, 6}}}, [][]int{{8, 2, 3}, {9, 6, 7}, {4, 5, 1}}},
		{"Example 2", args{grid: [][]int{{0, 1}, {1, 2}}}, [][]int{{2, 1}, {1, 0}}},
		{"Example 3", args{grid: [][]int{{1}}}, [][]int{{1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortMatrix(tt.args.grid); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
