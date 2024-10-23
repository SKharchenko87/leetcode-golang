package p2641

import (
	"leetcode/stucture"
	"reflect"
	"testing"
)

const null = stucture.NULL

func Test_run(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Example 1", args{[]int{5, 4, 9, 1, 10, null, 7}}, []int{0, 0, 0, 7, 7, null, 11}},
		{"Example 2", args{[]int{3, 1, 2}}, []int{0, 0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
