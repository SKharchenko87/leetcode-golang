package p2191

import (
	"reflect"
	"testing"
)

func Test_sortJumbled(t *testing.T) {
	type args struct {
		mapping []int
		nums    []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Example 1", args{mapping: []int{8, 9, 4, 0, 2, 1, 3, 5, 7, 6}, nums: []int{991, 338, 38}}, []int{338, 38, 991}},
		{"Example 2", args{mapping: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, nums: []int{789, 456, 123}}, []int{123, 456, 789}},
		{"TestCase 13", args{mapping: []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, nums: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}}, []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}},
		{"TestCase 29", args{mapping: []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, nums: []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}}, []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortJumbled(tt.args.mapping, tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortJumbled() = %v, want %v", got, tt.want)
			}
		})
	}
}
