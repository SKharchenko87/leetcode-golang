package p2415

import (
	"reflect"
	"testing"
)

//import . "leetcode/stucture"

func Test_run(t *testing.T) {
	type args struct {
		root []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Example 1", args{root: []int{2, 3, 5, 8, 13, 21, 34}}, []int{2, 5, 3, 8, 13, 21, 34}},
		{"Example 2", args{root: []int{7, 13, 11}}, []int{7, 11, 13}},
		{"Example 3", args{root: []int{0, 1, 2, 0, 0, 0, 0, 1, 1, 1, 1, 2, 2, 2, 2}}, []int{0, 2, 1, 0, 0, 0, 0, 2, 2, 2, 2, 1, 1, 1, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
