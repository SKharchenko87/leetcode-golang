package p2161

import (
	"reflect"
	"testing"
)

func Test_pivotArray(t *testing.T) {
	type args struct {
		nums  []int
		pivot int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Example 1", args{nums: []int{9, 12, 5, 10, 14, 3, 10}, pivot: 10}, []int{9, 5, 3, 10, 10, 12, 14}},
		{"Example 2", args{nums: []int{-3, 4, 3, 2}, pivot: 2}, []int{-3, 2, 4, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pivotArray(tt.args.nums, tt.args.pivot); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pivotArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
