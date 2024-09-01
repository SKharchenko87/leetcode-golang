package p2022

import (
	"reflect"
	"testing"
)

func Test_construct2DArray(t *testing.T) {
	type args struct {
		original []int
		m        int
		n        int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"Example 1", args{original: []int{1, 2, 3, 4}, m: 2, n: 2}, [][]int{{1, 2}, {3, 4}}},
		{"Example 2", args{original: []int{1, 2, 3}, m: 1, n: 3}, [][]int{{1, 2, 3}}},
		{"Example 3", args{original: []int{1, 2}, m: 1, n: 1}, [][]int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := construct2DArray(tt.args.original, tt.args.m, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("construct2DArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
