package p1380

import (
	"reflect"
	"testing"
)

func Test_luckyNumbers(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Example 1", args{matrix: [][]int{{3, 7, 8}, {9, 11, 13}, {15, 16, 17}}}, []int{15}},
		{"Example 2", args{matrix: [][]int{{1, 10, 4, 2}, {9, 3, 8, 7}, {15, 16, 17, 12}}}, []int{12}},
		{"Example 3", args{matrix: [][]int{{7, 8}, {1, 2}}}, []int{7}},
		{"My 1", args{matrix: [][]int{{1, 8}, {0, 2}}}, []int{1}},
		{"My 2", args{matrix: [][]int{{1, 7}, {8, 2}}}, []int{}},
		{"Testcase 102", args{matrix: [][]int{{4162}, {65091}}}, []int{65091}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := luckyNumbers(tt.args.matrix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("luckyNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
