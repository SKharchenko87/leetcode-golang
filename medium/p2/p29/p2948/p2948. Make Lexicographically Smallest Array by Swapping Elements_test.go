package p2948

import (
	"reflect"
	"testing"
)

func Test_lexicographicallySmallestArray(t *testing.T) {
	type args struct {
		nums  []int
		limit int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Example 1", args{nums: []int{1, 5, 3, 9, 8}, limit: 2}, []int{1, 3, 5, 8, 9}},
		{"Example 2", args{nums: []int{1, 7, 6, 18, 2, 1}, limit: 3}, []int{1, 6, 7, 18, 1, 2}},
		{"Example 3", args{nums: []int{1, 7, 28, 19, 10}, limit: 3}, []int{1, 7, 28, 19, 10}},
		{"TestCase 277", args{nums: []int{1, 60, 34, 84, 62, 56, 39, 76, 49, 38}, limit: 4}, []int{1, 56, 34, 84, 60, 62, 38, 76, 49, 39}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lexicographicallySmallestArray(tt.args.nums, tt.args.limit); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("lexicographicallySmallestArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
