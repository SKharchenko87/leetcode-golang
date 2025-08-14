package p1038

import (
	"reflect"
	"testing"
)
import . "leetcode/stucture"

const null = NULL

func Test_helper(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Example 1", args{[]int{4, 1, 6, 0, 2, 5, 7, null, null, null, 3, null, null, null, 8}}, []int{30, 36, 21, 36, 35, 26, 15, null, null, null, 33, null, null, null, 8}},
		{"Example 2", args{[]int{0, null, 1}}, []int{1, null, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := helper(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("helper() = %v, want %v", got, tt.want)
			}
		})
	}
}
