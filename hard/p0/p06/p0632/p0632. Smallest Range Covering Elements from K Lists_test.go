package p0632

import (
	"reflect"
	"testing"
)

func Test_smallestRange(t *testing.T) {
	type args struct {
		nums [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Example 1", args{nums: [][]int{{4, 10, 15, 24, 26}, {0, 9, 12, 20}, {5, 18, 22, 30}}}, []int{20, 24}},
		{"Example 2", args{nums: [][]int{{1, 2, 3}, {1, 2, 3}, {1, 2, 3}}}, []int{1, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := smallestRange(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("smallestRange() = %v, want %v", got, tt.want)
			}
		})
	}
}
