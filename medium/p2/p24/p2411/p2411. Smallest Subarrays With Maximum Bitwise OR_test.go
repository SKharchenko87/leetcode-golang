package p2411

import (
	"reflect"
	"testing"
)

func Test_smallestSubarrays(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Example 1", args{nums: []int{1, 0, 2, 1, 3}}, []int{3, 3, 2, 2, 1}},
		{"Example 2", args{nums: []int{1, 2}}, []int{2, 1}},
		{"TestCase 20", args{nums: []int{9, 5, 0, 10, 7, 2, 9, 2, 4}}, []int{4, 3, 3, 2, 3, 4, 3, 2, 1}},
		{"TestCase 33", args{nums: []int{4, 0, 5, 6, 3, 2}}, []int{4, 3, 2, 2, 1, 1}},
		{"TestCase 46", args{nums: []int{1, 0}}, []int{1, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := smallestSubarrays(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("smallestSubarrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
