package p0498

import (
	"reflect"
	"testing"
)

func Test_findDiagonalOrder(t *testing.T) {
	type args struct {
		mat [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Example 1", args{mat: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}}, []int{1, 2, 4, 7, 5, 3, 6, 8, 9}},
		{"Example 2", args{mat: [][]int{{1, 2}, {3, 4}}}, []int{1, 2, 3, 4}},
		{"TestCase 14", args{mat: [][]int{{2, 5}, {8, 4}, {0, -1}}}, []int{2, 5, 8, 0, 4, -1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDiagonalOrder(tt.args.mat); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findDiagonalOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
