package p1171

import (
	"reflect"
	"testing"
)

func Test_run(t *testing.T) {
	type args struct {
		head []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Example 1", args{[]int{1, 2, -3, 3, 1}}, []int{3, 1}},
		{"Example 2", args{[]int{1, 2, 3, -3, 4}}, []int{1, 2, 4}},
		{"Example 3", args{[]int{1, 2, 3, -3, -2}}, []int{1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
