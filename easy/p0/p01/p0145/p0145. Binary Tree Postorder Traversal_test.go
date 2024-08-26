package p0145

import (
	. "leetcode/stucture"
	"reflect"
	"testing"
)

const null = NULL

func Test_run(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Example 1", args{[]int{1, null, 2, 3}}, []int{3, 2, 1}},
		{"Example 2", args{[]int{}}, []int{}},
		{"Example 3", args{[]int{1}}, []int{1}},
		{"My 1", args{[]int{1, 4, 2, 3, 5, null, 6}}, []int{3, 5, 4, 6, 2, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
