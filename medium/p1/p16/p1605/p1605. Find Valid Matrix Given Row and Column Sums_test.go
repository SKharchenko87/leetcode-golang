package p1605

import (
	"reflect"
	"testing"
)

func Test_restoreMatrix(t *testing.T) {
	type args struct {
		rowSum []int
		colSum []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"Example 1", args{rowSum: []int{3, 8}, colSum: []int{4, 7}}, [][]int{{3, 0}, {1, 7}}},
		{"Example 2", args{rowSum: []int{5, 7, 10}, colSum: []int{8, 6, 8}}, [][]int{{5, 0, 0}, {3, 4, 0}, {0, 2, 8}}},
		{"Testcase 2", args{rowSum: []int{14, 9}, colSum: []int{6, 9, 8}}, [][]int{{6, 8, 0}, {0, 1, 8}}},
		{"Testcase 81", args{rowSum: []int{1, 0}, colSum: []int{1}}, [][]int{{1}, {0}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := restoreMatrix(tt.args.rowSum, tt.args.colSum); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("restoreMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
