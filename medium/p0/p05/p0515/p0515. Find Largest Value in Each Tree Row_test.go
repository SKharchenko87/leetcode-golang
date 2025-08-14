package p0515

import (
	"reflect"
	"testing"
)

import . "leetcode/stucture"

func Test_run(t *testing.T) {
	type args struct {
		root []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Example 1", args{root: []int{1, 3, 2, 5, 3, NULL, 9}}, []int{1, 3, 9}},
		{"Example 2", args{root: []int{1, 2, 3}}, []int{1, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
