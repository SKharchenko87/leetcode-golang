package p2181

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
		{"Example 1", args{[]int{0, 3, 1, 0, 4, 5, 2, 0}}, []int{4, 11}},
		{"Example 2", args{[]int{0, 1, 0, 3, 0, 2, 2, 0}}, []int{1, 3, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
