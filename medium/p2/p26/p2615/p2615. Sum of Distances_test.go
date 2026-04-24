package p2615

import (
	"reflect"
	"testing"
)

func Test_distance(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		{"Example 1", args{nums: []int{1, 3, 1, 1, 2}}, []int64{5, 0, 3, 4, 0}},
		{"Example 2", args{nums: []int{0, 5, 3}}, []int64{0, 0, 0}},
		{"My 1", args{nums: []int{1, 3, 1, 1, 2, 1, 1}}, []int64{16, 0, 10, 9, 0, 11, 14}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := distance(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("distance() = %v, want %v", got, tt.want)
			}
		})
	}
}
