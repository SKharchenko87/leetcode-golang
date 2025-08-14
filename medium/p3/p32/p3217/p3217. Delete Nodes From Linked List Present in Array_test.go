package p3217

import (
	"reflect"
	"testing"
)

func Test_run(t *testing.T) {
	type args struct {
		nums []int
		root []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Example 1", args{nums: []int{1, 2, 3}, root: []int{1, 2, 3, 4, 5}}, []int{4, 5}},
		{"Example 2", args{nums: []int{1}, root: []int{1, 2, 1, 2, 1, 2}}, []int{2, 2, 2}},
		{"Example 3", args{nums: []int{5}, root: []int{1, 2, 3, 4}}, []int{1, 2, 3, 4}},
		{"TestCase 330", args{nums: []int{1, 7, 6, 2, 4}, root: []int{3, 7, 1, 8, 1}}, []int{3, 8}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.nums, tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
