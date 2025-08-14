package p2058

import (
	"reflect"
	"testing"
)

func Test_run(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Example 1", args{[]int{3, 1}}, []int{-1, -1}},
		{"Example 2", args{[]int{5, 3, 1, 2, 5, 1, 2}}, []int{1, 3}},
		{"Example 3", args{[]int{1, 3, 2, 2, 3, 2, 2, 2, 7}}, []int{3, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
