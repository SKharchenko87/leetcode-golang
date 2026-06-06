package p2574

import (
	"reflect"
	"testing"
)

func Test_leftRightDifference(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Example 1", args{nums: []int{10, 4, 8, 3}}, []int{15, 1, 11, 22}},
		{"Example 2", args{nums: []int{1}}, []int{0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := leftRightDifference(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("leftRightDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
