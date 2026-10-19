package p2089

import (
	"reflect"
	"testing"
)

func Test_targetIndices(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Example 1", args{nums: []int{1, 2, 5, 2, 3}, target: 2}, []int{1, 2}},
		{"Example 2", args{nums: []int{1, 2, 5, 2, 3}, target: 3}, []int{3}},
		{"Example 3", args{nums: []int{1, 2, 5, 2, 3}, target: 5}, []int{4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := targetIndices(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("targetIndices() = %v, want %v", got, tt.want)
			}
		})
	}
}
