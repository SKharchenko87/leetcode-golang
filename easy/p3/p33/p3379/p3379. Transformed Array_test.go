package p3379

import (
	"reflect"
	"testing"
)

func Test_constructTransformedArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Example 1", args{nums: []int{3, -2, 1, 1}}, []int{1, 1, 1, 3}},
		{"Example 2", args{nums: []int{-1, 4, -1}}, []int{-1, -1, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := constructTransformedArray(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("constructTransformedArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
