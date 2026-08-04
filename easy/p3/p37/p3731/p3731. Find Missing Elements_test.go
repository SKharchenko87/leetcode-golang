package p3731

import (
	"reflect"
	"testing"
)

func Test_findMissingElements(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Example 1", args{nums: []int{1, 4, 2, 5}}, []int{3}},
		{"Example 2", args{nums: []int{7, 8, 6, 9}}, []int{}},
		{"Example 3", args{nums: []int{5, 1}}, []int{2, 3, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMissingElements(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findMissingElements() = %v, want %v", got, tt.want)
			}
		})
	}
}
