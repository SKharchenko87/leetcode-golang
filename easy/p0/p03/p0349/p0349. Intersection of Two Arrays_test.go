package p0349

import (
	"reflect"
	"testing"
)

func Test_intersection(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Example 1", args{[]int{1, 2, 2, 1}, []int{2, 2}}, []int{2}},
		{"Example 2", args{[]int{4, 9, 5}, []int{9, 4, 9, 8, 4}}, []int{4, 9}},
		{"Example 5", args{[]int{1, 2}, []int{1, 1}}, []int{1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intersection(tt.args.nums1, tt.args.nums2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("intersection() = %v, want %v", got, tt.want)
			}
		})
	}
}
