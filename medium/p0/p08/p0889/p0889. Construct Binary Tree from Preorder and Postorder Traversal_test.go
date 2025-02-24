package p0889

import (
	"reflect"
	"testing"
)

func Test_run(t *testing.T) {
	type args struct {
		preorder  []int
		postorder []int
	}
	tests := []struct {
		name string
		args args
		want []any
	}{
		{"Example 1", args{preorder: []int{1, 2, 4, 5, 3, 6, 7}, postorder: []int{4, 5, 2, 6, 7, 3, 1}}, []any{1, 2, 3, 4, 5, 6, 7}},
		{"Example 2", args{preorder: []int{1}, postorder: []int{1}}, []any{1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.preorder, tt.args.postorder); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
