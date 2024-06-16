package p0075

import (
	"reflect"
	"testing"
)

func Test_run(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{

		{"Example 1", args{[]int{2, 0, 2, 1, 1, 0}}, []int{0, 0, 1, 1, 2, 2}},
		{"Example 2", args{[]int{2, 0, 1}}, []int{0, 1, 2}},
		{"My 1", args{[]int{1, 0, 1, 0, 2, 2}}, []int{0, 0, 1, 1, 2, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
